package policy

import "github.com/jb843051627/mireflux/internal/model"

func FieldDiagnostics(readings []model.Reading) []model.FieldDiagnostic {
	return []model.FieldDiagnostic{
		CO2BaselineDiagnostic(readings),
		CO2StepDiagnostic(readings),
		CO2RateDiagnostic(readings),
		CO2TemperatureCouplingDiagnostic(readings),
		CO2PressureCouplingDiagnostic(readings),
		TemperatureBaselineDiagnostic(readings),
		TemperatureStepDiagnostic(readings),
		TemperatureRateDiagnostic(readings),
		TemperatureHumidityMarginDiagnostic(readings),
		PressureBaselineDiagnostic(readings),
		PressureStepDiagnostic(readings),
		PressureRateDiagnostic(readings),
		PressureHumidityCouplingDiagnostic(readings),
		HumidityBaselineDiagnostic(readings),
		HumidityStepDiagnostic(readings),
		HumidityRateDiagnostic(readings),
		HumidityTemperatureCouplingDiagnostic(readings),
		ArrivalLatencyDiagnostic(readings),
		SamplingGapDiagnostic(readings),
		CollectionHourDiagnostic(readings),
		CollectionMinuteDiagnostic(readings),
		LabelCoverageDiagnostic(readings),
		ChamberIdentifierDiagnostic(readings),
		ReceiveOrderDiagnostic(readings),
		CO2RangeDiagnostic(readings),
		ThermalRangeDiagnostic(readings),
		PressureRangeDiagnostic(readings),
		HumidityRangeDiagnostic(readings),
		HeadspaceDensityDiagnostic(readings),
		DewMarginDiagnostic(readings),
		ThermalPressureProductDiagnostic(readings),
		CollectionSequenceDiagnostic(readings),
		CO2ReferenceMarginDiagnostic(readings),
		CO2FluxEnvelopeDiagnostic(readings),
		CO2HumidityCouplingDiagnostic(readings),
		CO2ThermalProductDiagnostic(readings),
		TemperatureComfortDiagnostic(readings),
		TemperatureExcursionDiagnostic(readings),
		TemperaturePressureRatioDiagnostic(readings),
		TemperatureCollectionWindowDiagnostic(readings),
		PressureComfortDiagnostic(readings),
		PressureExcursionDiagnostic(readings),
		PressureThermalCouplingDiagnostic(readings),
		PressureCollectionWindowDiagnostic(readings),
		HumidityComfortDiagnostic(readings),
		HumidityExcursionDiagnostic(readings),
		HumidityPressureProductDiagnostic(readings),
		HumidityCollectionWindowDiagnostic(readings),
		CycleSpacingDiagnostic(readings),
		ReceptionLagDiagnostic(readings),
		LabelRichnessDiagnostic(readings),
		ChamberCodeLengthDiagnostic(readings),
		FieldClockHourDiagnostic(readings),
		FieldClockMinuteDiagnostic(readings),
		SampleOrdinalDiagnostic(readings),
		HeadspaceTemperatureProductDiagnostic(readings),
		AtmosphericMoistureIndexDiagnostic(readings),
		CO2TimingDiagnostic(readings),
		CO2WeatherOffsetDiagnostic(readings),
		ThermalMedianDiagnostic(readings),
		ThermalDewSpanDiagnostic(readings),
		PressureContinuityDiagnostic(readings),
		HumidityContinuityDiagnostic(readings),
		RecordingLatencyBudgetDiagnostic(readings),
		RadioCadenceDiagnostic(readings),
		LabelPresenceDiagnostic(readings),
		LabelPresenceDiagnostic(readings),
		TemporalHorizonDiagnostic(readings),
		ProgressiveOrderDiagnostic(readings),
	}
}

func DiagnosticScore(checks []model.FieldDiagnostic) float64 {
	if len(checks) == 0 {
		return 0
	}
	score := 1.0
	for _, check := range checks {
		switch check.State {
		case model.DiagnosticWatch:
			score -= 0.015
		case model.DiagnosticBlocker:
			score -= 0.060
		}
	}
	if score < 0 {
		return 0
	}
	return score
}
