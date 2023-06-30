// the logger package is used to log messages to the console
package logger

import (
	"github.com/markot99/myinventory-backend/utils/envutils"
	"github.com/sirupsen/logrus"
)

// InitializeLogger is used to initialize the logger
func InitializeLogger() {
	logLevelString := envutils.GetEnvVariableOrDefault("LOG_LEVEL", "info")
	level, err := logrus.ParseLevel(logLevelString)
	if err != nil {
		logrus.Error("Failed to parse log level: '", logLevelString, "'. Using default level: info")
	}
	logrus.SetLevel(level)
}

// Debug is used to log debug messages
func Debug(args ...interface{}) {
	logrus.Debug(args...)
}

// Info is used to log info messages
func Info(args ...interface{}) {
	logrus.Info(args...)
}

// Warn is used to log warning messages
func Warn(args ...interface{}) {
	logrus.Warn(args...)
}

// Error is used to log error messages
func Error(args ...interface{}) {
	logrus.Error(args...)
}

// Fatal is used to log fatal messages
func Fatal(args ...interface{}) {
	logrus.Fatal(args...)
}
