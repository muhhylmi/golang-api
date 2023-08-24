package utils

import (
	"github.com/sirupsen/logrus"
)

func Newlogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetLevel(logrus.DebugLevel)

	return logger
}

func LogWithContext(logger *logrus.Logger, ctx string, scope string) *logrus.Entry {
	return logger.WithFields(logrus.Fields{
		"scope":   scope,
		"context": ctx,
	})
}
