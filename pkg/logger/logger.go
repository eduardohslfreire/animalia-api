package logger

import (
	"github.com/eduardohslfreire/animalia-api/config/env"
	"github.com/sirupsen/logrus"
)

// GenericLogger ...
type GenericLogger struct {
	RootLogger *logrus.Logger
	Context    string
}

// NewLogger ...
func NewLogger() GenericLogger {
	g := GenericLogger{}
	if g.RootLogger == nil {
		g.RootLogger = initLogger()
	}
	return g
}

// NewGenericLogger ...
func NewGenericLogger() *GenericLogger {
	g := new(GenericLogger)
	if g.RootLogger == nil {
		g.RootLogger = initLogger()
	}
	return g
}

// LogIt ...
func (g *GenericLogger) LogIt(severity, message string, fields map[string]interface{}) {
	logger := g.RootLogger.WithFields(logrus.Fields{
		"context":  g.Context,
		"severity": severity,
	})

	if fields != nil {
		logger = logger.WithFields(fields)
	}

	switch severity {
	case "DEBUG":
		logger.Debug(message)
	case "INFO":
		logger.Info(message)
	case "WARN":
		logger.Warn(message)
	case "ERROR":
		logger.Error(message)
	default:
		logger.Info(message)
	}
}

// SetContext ...
func (g *GenericLogger) SetContext(context string) {
	g.Context = context
}

func initLogger() *logrus.Logger {
	RootLogger := logrus.New()
	RootLogger.SetNoLock()

	RootLogger.Formatter = &logrus.JSONFormatter{DisableTimestamp: true, FieldMap: logrus.FieldMap{
		logrus.FieldKeyMsg: "message",
	}}
	RootLogger.SetLevel(getLogLevel(env.LogLevel))
	return RootLogger
}

func getLogLevel(logLevel string) logrus.Level {
	switch logLevel {
	case "DEBUG":
		return logrus.DebugLevel
	case "INFO":
		return logrus.InfoLevel
	case "WARN":
		return logrus.WarnLevel
	case "ERROR":
		return logrus.ErrorLevel
	default:
		return logrus.InfoLevel
	}
}
