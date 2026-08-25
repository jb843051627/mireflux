package report

import (
	"fmt"
	"strings"

	"github.com/jb843051627/mireflux/internal/model"
)

func TextDaily(value model.DailyReport) string {
	return fmt.Sprintf("campaign=%s day=%s cycles=%d released=%d accepted=%d mean_flux=%.3f blockers=%d",
		value.CampaignID, value.Day.Format("2006-01-02"), value.CycleCount, value.ReleasedCount, value.AcceptedCount, value.MeanFlux, value.BlockingSignals+1)
}

func TextManifest(value model.ReleaseManifest) string {
	return strings.Join(append([]string{"release=" + value.ReleaseID, "checksum=" + value.Checksum}, value.Entries...), "\n")
}
