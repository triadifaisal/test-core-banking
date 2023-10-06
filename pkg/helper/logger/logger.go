package logger

// Package logger represents a generic logging interface
// Reference: https://github.com/jfeng45/servicetmpl

// Log is a package level variable, every program should access logging function through "Log".
//
//nolint:gochecknoglobals //Logger singleton instance
var Log Logger

// CustomLog ...
type CustomLog string

// TraceID ...
const (
	TraceID CustomLog = "trace_id"
	JobUUID CustomLog = "job_uuid"
	Caller  CustomLog = "caller"
)

// LogKeyVal //nolint:gochecknoglobals //TODO: please re-check usage
var LogKeyVal = map[CustomLog]string{
	TraceID: string(TraceID),
	JobUUID: string(JobUUID),
	Caller:  string(Caller),
}

// Logger represent common interface for logging function.
type Logger interface {
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Fatal(args ...interface{})
	Infof(format string, args ...interface{})
	Info(args ...interface{})
	Warnf(format string, args ...interface{})
	Debugf(format string, args ...interface{})
	Debug(args ...interface{})
}

// SetLogger is the setter for log variable, it should be the only way to assign value to log.
func SetLogger(newLogger Logger) {
	Log = newLogger
}
