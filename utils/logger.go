package utils

import (
	"github.com/sirupsen/logrus"
)

func Newlogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:            false, // Nonaktifkan warna
		DisableColors:          true,  // Matikan warna
		DisableTimestamp:       true,  // Nonaktifkan tanggal
		DisableLevelTruncation: true,  // Nonaktifkan pemotongan level
		FullTimestamp:          true,  // Tampilkan timestamp lengkap (jika DisableTimestamp=false)
	})
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
