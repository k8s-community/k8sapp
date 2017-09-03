package logger

import "github.com/sirupsen/logrus"

// newLogrus creates "github.com/sirupsen/logrus" logger
func newLogrus(config *Config) Logger {
	logger := logrus.New()
	logger.Level = logrusLevelConverter(config.Level)
	logger.WithFields(logrus.Fields(config.Fields))
	return logger
}

func logrusLevelConverter(level Level) logrus.Level {
	switch level {
	case LevelDebug:
		return logrus.DebugLevel
	case LevelInfo:
		return logrus.InfoLevel
	case LevelWarn:
		return logrus.WarnLevel
	case LevelError:
		return logrus.ErrorLevel
	case LevelFatal:
		return logrus.FatalLevel
	default:
		return logrus.InfoLevel
	}
}
