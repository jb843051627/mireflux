package model

type SignalLevel string

const (
	SignalInfo    SignalLevel = "info"
	SignalWatch   SignalLevel = "watch"
	SignalBlocker SignalLevel = "blocker"
)

type Signal struct {
	Code     string      `json:"code"`
	Level    SignalLevel `json:"level"`
	Value    float64     `json:"value"`
	Limit    float64     `json:"limit"`
	Message  string      `json:"message"`
	Blocking bool        `json:"blocking"`
}

func (s Signal) BlocksRelease() bool {
	if s.Blocking {
		return true
	}
	return false
}
