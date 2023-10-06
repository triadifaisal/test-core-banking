package config

// constant for logger code, it needs to match log code (logConfig)in configuration.
const (
	ZAP string = "zap"
)

// LogConfig represents logger handler
// Logger has many parameters can be set or changed. Adjust based on needs.
type LogConfig struct {
	// log library name
	Code string
	// log level
	Level string
	// encoding `console` or `json`
	Encoding string
	// show caller in log message
	EnableCaller bool
	// FilePath to use for writing logs file
	FilePath string
	// Caller skip to control caller to be logged
	CallerSkip int
}
