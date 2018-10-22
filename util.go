package log

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	tContext "gitlab.mytaxi.lk/pickme/go-util/traceable_context"
	"log"
	"runtime"
)

type logParser struct {
	*logOptions
	log *log.Logger
}

//isLoggable Check whether the log type is loggable under current configurations
func (l *logParser) isLoggable(logType string) bool {
	return logTypes[logType] <= logTypes[string(l.logLevel)]
}

func (l *logParser) colored(typ string) string {
	if l.colors {
		return logColors[typ]
	}

	return typ
}

func (l *logParser) toString(id string, typ string, message interface{}, params ...interface{}) string {

	var messageFmt = "%s %s %v"

	return fmt.Sprintf(messageFmt,
		typ,
		fmt.Sprintf("%+v", message),
		fmt.Sprintf("%+v", params))
}

func (l *logParser) logEntryContext(logType string, ctx context.Context, message interface{}, color string, params ...interface{}) {
	l.logEntry(logType, uuidFromContext(ctx), message, color, params...)
}

func WithPrefix(p string, message interface{}) string {
	return fmt.Sprintf(`%s] [%+v`, p, message)
}

func uuidFromContext(ctx context.Context) uuid.UUID {
	traceableCtx, ok := ctx.(tContext.TraceableContext)
	if !ok {
		return uuid.New()
	}
	return traceableCtx.UUID()
}

func (l *logParser) logEntry(logType string, uuid uuid.UUID, message interface{}, color string, params ...interface{}) {

	if !l.isLoggable(logType) {
		return
	}

	var file string
	var line int
	if l.filePath {
		_, f, l, ok := runtime.Caller(l.fileDepth)
		if !ok {
			f = `<Unknown>`
			l = 1
		}

		file = f
		line = l

		message = fmt.Sprintf(`[%s] [%+v on %s %d]`, uuid.String(), message, file, line)
	} else {
		message = fmt.Sprintf(`[%s] [%+v]`, uuid.String(), message)
	}

	if logType == fatal {
		l.log.Fatalln(l.toString(``, color, message, params...))
	}

	if logType == err {
		l.log.Println(l.toString(``, color, message, params...))
		return
	}

	l.log.Println(l.toString(``, color, message, params...))
}