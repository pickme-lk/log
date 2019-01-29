package log

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	tContext "gitlab.mytaxi.lk/pickme/go-util/traceable_context"
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

func (l *logParser) toString(logMsg *logMessage, params ...interface{}) string {
	if len(params) > 0 {
		return fmt.Sprintf("%s [%s] [%+v on %s line %d] %+v", logMsg.typ, logMsg.uuid, logMsg.message, logMsg.file, logMsg.line, params)
	}
	return fmt.Sprintf("%s [%s] [%+v on %s line %d]", logMsg.typ, logMsg.uuid, logMsg.message, logMsg.file, logMsg.line)
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

func (l *logParser) logEntry(level Level, ctx context.Context, message interface{}, params ...interface{}) {

	if !l.isLoggable(level) {
		return
	}

	var uid uuid.UUID

	if ctx != nil {
		uid = uuidFromContext(ctx)
	} else {
		uid = uuid.New()
	}

	logLevel := string(level)

	if l.colors {
		logLevel = l.colored(level)
	}

	logMsg := &logMessage{
		typ:     logLevel,
		message: message,
		uuid:    uid.String(),
	}

	if l.filePath {
		_, f, l, ok := runtime.Caller(l.fileDepth)
		if !ok {
			f = `<Unknown>`
			l = 1
		}

		logMsg.file = f
		logMsg.line = l

		//message = fmt.Sprintf(`[%s] [%+v on %s %d]`, uid.String(), message, file, line)

	} else {
		//message = fmt.Sprintf(`[%s] [%+v]`, uid.String(), message)
	}

	if level == FATAL {
		l.log.Fatalln(l.toString(logMsg, params...))
	}

	l.log.Println(l.toString(logMsg, params...))
}
