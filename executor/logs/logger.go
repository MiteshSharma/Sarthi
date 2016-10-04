package logs

import (
	"github.com/Sirupsen/logrus"
)

type ExecutorLogger struct {
	Name string
	Log  *logrus.Logger
}

var Logger *ExecutorLogger

func NewExecutorLogger(name string) *ExecutorLogger {
	var logger = ExecutorLogger{Name: name}
	logger.Log = logrus.New()
	logger.Log.Level = logrus.DebugLevel
	return &logger
}

func (l *ExecutorLogger) Debug(msg string) {
	l.Log.Debug(msg)
}

func (l *ExecutorLogger) Info(msg string) {
	l.Log.Info(msg)
}

func (l *ExecutorLogger) Error(msg string) {
	l.Log.Error(msg)
}
