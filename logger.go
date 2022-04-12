package logger

import (
	"encoding/json"
	"log"
)

type Severity string

const (
	SeverityDefault   Severity = "DEFAULT"
	SeverityDebug     Severity = "DEBUG"
	SeverityInfo      Severity = "INFO"
	SeverityNotice    Severity = "NOTICE"
	SeverityWarning   Severity = "WARNING"
	SeverityError     Severity = "ERROR"
	SeverityCritical  Severity = "CRITICAL"
	SeverityAlert     Severity = "ALERT"
	SeverityEmergency Severity = "EMERGENCY"
)

// Entry defines a log entry.
type Entry struct {
	Message  string   `json:"message"`
	Severity Severity `json:"severity,omitempty"`
	Trace    string   `json:"logging.googleapis.com/trace,omitempty"`

	// Logs Explorer allows filtering and display of this as `jsonPayload.component`.
	Component string `json:"component,omitempty"`
}

func init() {
	// Disable log prefixes such as the default timestamp.
	// Prefix text prevents the message from being parsed as JSON.
	// A timestamp suppose to be added when shipping logs to Cloud Logging.
	log.SetFlags(0)
}

// String renders an entry structure to the JSON format expected by Cloud Logging.
func (e Entry) String() string {
	if e.Severity == "" {
		e.Severity = SeverityInfo
	}
	out, err := json.Marshal(e)
	if err != nil {
		log.Printf("json.Marshal: %v", err)
	}
	return string(out)
}

func Default(message string) {
	log.Println(Entry{
		Severity: SeverityDefault,
		Message:  message,
	})
}

func Debug(message string) {
	log.Println(Entry{
		Severity: SeverityDebug,
		Message:  message,
	})
}

func Info(message string) {
	log.Println(Entry{
		Severity: SeverityInfo,
		Message:  message,
	})
}

func Notice(message string) {
	log.Println(Entry{
		Severity: SeverityNotice,
		Message:  message,
	})
}

func Warning(message string) {
	log.Println(Entry{
		Severity: SeverityWarning,
		Message:  message,
	})
}

func Error(message string) {
	log.Println(Entry{
		Severity: SeverityError,
		Message:  message,
	})
}

func Critical(message string) {
	log.Println(Entry{
		Severity: SeverityCritical,
		Message:  message,
	})
}

func Alert(message string) {
	log.Println(Entry{
		Severity: SeverityAlert,
		Message:  message,
	})
}

func Emergency(message string) {
	log.Println(Entry{
		Severity: SeverityEmergency,
		Message:  message,
	})
}
