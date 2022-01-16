package util

import (
	"github.com/rs/zerolog"
)

const pathToLogConfig = "configuration/livesettings.json"
const pathToConfig = "configuration/constsettings.json"
const logLevel = "Logging.LogLevel.Debug"

func setLogLevel(level string) {
	switch level {
	case "Debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		break
	case "Info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		break
	case "Warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
		break
	case "Error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
		break
	default:
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	}
}
