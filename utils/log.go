package utils

import (
	"github.com/sirupsen/logrus"
	"os"
)

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.InfoLevel)
}

func LogInfo(s string) {
	logrus.Info(s)
}

func LogWarn(s string) {
	logrus.Warn(s)
}

func LogError(s string) {
	logrus.Error(s)
}

func LogPanic(s string) {
	logrus.Panic(s)
}
