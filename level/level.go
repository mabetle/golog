// package github.com/mabetle/golog/level
package level

// define Level
type Level struct{}

// log level
const (
   LogError   = 0
   LogWarn	  = 1
   LogInfo    = 2
   LogDebug   = 3
   LogTrace   = 4
   LogLevel	  = 2
)


// Get level num from string
func (l Level) GetMappedLogLevel(level string) int {
	level = strings.ToUpper(level)
	switch level {
		case "ERROR": return LogError
		case "WARN": return LogWarn
		case "INFO": return LogInfo
		case "DEBUG": return LogDebug
		case "TRACE": return LogTrace
	}
	return LogDebug
}


