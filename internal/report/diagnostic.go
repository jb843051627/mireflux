package report

import (
	"fmt"
	"strings"

	"github.com/jb843051627/mireflux/internal/model"
)

func TextDiagnostics(value model.DiagnosticReport) string {
	rows := make([]string, 0, len(value.Checks)+1)
	rows = append(rows, fmt.Sprintf("cycle=%s chamber=%s score=%.3f", value.CycleID, value.ChamberID, value.Score))
	for _, check := range value.Checks {
		rows = append(rows, fmt.Sprintf("%s=%s value=%.3f%s", check.Code, check.State, check.Value, check.Unit))
	}
	return strings.Join(rows, "\n")
}

func AttentionDiagnostics(value model.DiagnosticReport) []model.FieldDiagnostic {
	items := make([]model.FieldDiagnostic, 0)
	for _, check := range value.Checks {
		if check.RequiresAttention() {
			items = append(items, check)
		}
	}
	return items
}
