package log

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	tContext "github.com/pickme-go/traceable-context"
	"log"
	"runtime"
)

type logMessage struct {
	typ     string
	color   string
	message interface{}
	uuid    string
	file    string
	line    int
}

type logParser struct {
	*logOptions
	log *log.Logger
}

//isLoggable Check whether the log type is loggable under current configurations
func (l *logParser) isLoggable(level Level) bool {
	return logTypes[level] <= logTypes[l.logLevel]
}

func (l *logParser) colored(level Level) string {
	if l.colors {
		return string(logColors[level])
	}

	return string(level)
}

func (l *logParser) WithPrefix(p string, message interface{}) string {
	if l.prefix != `` {
		return fmt.Sprintf(`%s%s] [%+v`, l.prefix, p, message)
	}
	return fmt.Sprintf(`%s] [%+v`, p, message)
}

func WithPrefix(p string, message interface{}) string {
	return fmt.Sprintf(`%s] [%+v`, p, message)
}

func uuidFromContext(ctx context.Context) uuid.UUID {
	uid := tContext.FromContext(ctx)
	if uid == uuid.Nil {
		return uuid.New()
	}

	return uid
}

func (l *logParser) logEntry(level Level, ctx context.Context, message interface{}, prms ...interface{}) {
	if !l.isLoggable(level) {
		return
	}

	format := "%s [%s] [%+v]"

	var params []interface{}

	logLevel := string(level)

	if l.colors {
		logLevel = l.colored(level)
	}

	var uid uuid.UUID
	if ctx != nil {
		uid = uuidFromContext(ctx)
	} else {
		uid = uuid.New()
	}

	logMsg := &logMessage{
		typ:     logLevel,
		message: message,
		uuid:    uid.String(),
	}

	params = append(params, logLevel, uid.String(), message)

	if l.filePath {
		_, f, l, ok := runtime.Caller(l.fileDepth)
		if !ok {
			f = `<Unknown>`
			l = 1
		}

		logMsg.file = f
		logMsg.line = l

		format = "%s [%s] [%+v on %s line %d]"

		params = append(params, f, l)

	}

	if len(prms) > 0 {
		format = "%s [%s] [%+v on %s line %d] %+v"
		params = append(params, prms)
	}

	if level == FATAL {
		l.log.Fatalf(format, params...)
	}

	l.log.Printf(format, params...)
}
