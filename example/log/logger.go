package log

import "github.com/pickme-go/log"

type Logger struct {
	log log.Logger
}

func NewLogger() Logger {
	return Logger{
		log: log.Constructor.Log(
			log.WithLevel(log.TRACE),
			log.FileDepth(3),
			log.WithFilePath(true),
			log.Prefixed(`other_log`),
		),
	}
}

func (l *Logger) Info(v ...interface{}) {
	l.log.Info(v)
}

func (l *Logger) Trace(v ...interface{}) {
	l.log.Trace(v)
}
