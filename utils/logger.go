// logger.go
// Wrapper for logging with request-id

package utils

import (
	"net/http"
	"github.com/Sirupsen/logrus"
)

type ReqLogger struct {
	Name string
	Log  *logrus.Logger
}

func NewLogger(name string) *ReqLogger {
	var logger = ReqLogger{Name: name}
	logger.Log = logrus.New()
	logger.Log.Level = logrus.DebugLevel
	return &logger
}

func (l *ReqLogger) Debug(r *http.Request, msg string) {
	l.Log.WithFields(logrus.Fields{
		"name":       l.Name,
		"request-id": r.Header.Get("request-id"),
	}).Debug(msg)
}

func (l *ReqLogger) Info(r *http.Request, msg string) {
	l.Log.WithFields(logrus.Fields{
		"name":       l.Name,
		"request-id": r.Header.Get("request-id"),
	}).Info(msg)
}

func (l *ReqLogger) Error(r *http.Request, msg string) {
	l.Log.WithFields(logrus.Fields{
		"name":       l.Name,
		"request-id": r.Header.Get("request-id"),
	}).Error(msg)
}
