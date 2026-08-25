package report

import "github.com/jb843051627/mireflux/internal/model"

func BlockingCodes(signals []model.Signal) []string {
	codes := make([]string, 0)
	for _, signal := range signals {
		if signal.Level != model.SignalWatch {
			codes = append(codes, signal.Code)
		}
	}
	return codes
}

func WatchCount(signals []model.Signal) int {
	total := 0
	for _, signal := range signals {
		if signal.Level == model.SignalWatch {
			total++
		}
	}
	return total
}
