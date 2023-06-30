package envutils

import (
	"os"

	"github.com/sirupsen/logrus"
)

// GetEnvVariableOrDefault is used to get an environment variable or return a default value
func GetEnvVariableOrDefault(variableName, defaultValue string) string {
	value := os.Getenv(variableName)
	if value == "" {
		value = defaultValue
		logrus.Debug("Environment variable for ", variableName, " not set, using default value: ", defaultValue)
	}
	return value
}
