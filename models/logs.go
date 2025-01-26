package models

import "github.com/google/uuid"

const (
	LogSourceAPI       = "API"
	LogSourceAlerts    = "Alerts"
	LogSourceInterface = "Interface"

	LogLevelNone  = "None"
	LogLevelInfo  = "Info"
	LogLevelWarn  = "Warn"
	LogLevelDebug = "Debug"
	LogLevelError = "Error"
	LogLevelFatal = "Fatal"
)

func IsValidLogSource(s string) bool {
	return s == LogSourceAPI || s == LogSourceAlerts || s == LogSourceInterface
}

func IsValidLogLevel(s string) bool {
	return s == LogLevelNone || s == LogLevelInfo || s == LogLevelWarn || s == LogLevelDebug || s == LogLevelError || s == LogLevelFatal
}

type LogEdit struct {
	ID           uuid.UUID
	UserID       *uuid.UUID
	Message      string
	Context      *string
	RequestBody  *string
	StackTrace   *string
	StatusCode   int32
	ApiUrl       string
	InterfaceUrl string
	Level        string
	Source       string
}
