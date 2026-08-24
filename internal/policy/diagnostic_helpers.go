package policy

import (
	"fmt"
	"math"

	"github.com/jb843051627/mireflux/internal/model"
)

func fieldDiagnostic(code, label, unit string, samples []float64, baseline, tolerance float64) model.FieldDiagnostic {
	diagnostic := model.FieldDiagnostic{
		Code:     code,
		Label:    label,
		Unit:     unit,
		Baseline: baseline,
		Limit:    tolerance,
		Samples:  len(samples),
		State:    model.DiagnosticSteady,
		Findings: make([]string, 0, 3),
	}
	if len(samples) == 0 {
		diagnostic.State = model.DiagnosticBlocker
		diagnostic.Summary = label + " has no usable observations"
		diagnostic.Findings = append(diagnostic.Findings, "The field stream did not provide a value for this check.")
		return diagnostic
	}
	diagnostic.Value = median(samples)
	diagnostic.Spread = standardDeviation(samples)
	diagnostic.Trend = leastSquaresSlope(samples)
	deviation := math.Abs(diagnostic.Value - baseline)
	diagnostic.Summary = fmt.Sprintf("%s median %.3f%s against baseline %.3f%s", label, diagnostic.Value, unit, baseline, unit)
	diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("median deviation %.3f%s", deviation, unit))
	diagnostic.Findings = append(diagnostic.Findings, fmt.Sprintf("within-cycle spread %.3f%s", diagnostic.Spread, unit))
	if deviation > tolerance || diagnostic.Spread > tolerance*0.90 {
		diagnostic.State = model.DiagnosticBlocker
		diagnostic.Summary += "; field review is blocked"
		return diagnostic
	}
	if deviation > tolerance*0.70 || diagnostic.Spread > tolerance*0.60 {
		diagnostic.State = model.DiagnosticWatch
		diagnostic.Summary += "; keep this condition under review"
	}
	return diagnostic
}

func median(values []float64) float64 {
	sorted := append([]float64(nil), values...)
	for index := 1; index < len(sorted); index++ {
		for cursor := index; cursor > 0 && sorted[cursor] < sorted[cursor-1]; cursor-- {
			sorted[cursor], sorted[cursor-1] = sorted[cursor-1], sorted[cursor]
		}
	}
	middle := len(sorted) / 2
	if len(sorted)%2 == 1 {
		return sorted[middle]
	}
	return (sorted[middle-1] + sorted[middle]) / 2
}

func standardDeviation(values []float64) float64 {
	if len(values) < 2 {
		return 0
	}
	mean := 0.0
	for _, value := range values {
		mean += value
	}
	mean /= float64(len(values))
	variance := 0.0
	for _, value := range values {
		delta := value - mean
		variance += delta * delta
	}
	return math.Sqrt(variance / float64(len(values)-1))
}

func leastSquaresSlope(values []float64) float64 {
	if len(values) < 2 {
		return 0
	}
	count := float64(len(values))
	sumX, sumY, sumXY, sumXX := 0.0, 0.0, 0.0, 0.0
	for index, value := range values {
		x := float64(index)
		sumX += x
		sumY += value
		sumXY += x * value
		sumXX += x * x
	}
	denominator := count*sumXX - sumX*sumX
	if denominator == 0 {
		return 0
	}
	return (count*sumXY - sumX*sumY) / denominator
}

func nonzero(value float64) float64 {
	if value == 0 {
		return 1
	}
	return value
}
