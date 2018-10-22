package log

import (
	"context"
	"fmt"
	"github.com/google/uuid"
)

type prefixedLogger struct {
	logParser
}

func (l *prefixedLogger) ErrorContext(prefix string, ctx context.Context, message interface{}, params ...interface{}) {
	l.logEntryContext(err, ctx, WithPrefix(prefix, message), l.colored(`ERROR`), params...)
}

func (l *prefixedLogger) WarnContext(prefix string, ctx context.Context, message interface{}, params ...interface{}) {
	l.logEntryContext(warn, ctx, WithPrefix(prefix, message), l.colored(`WARN`), params...)
}

func (l *prefixedLogger) InfoContext(prefix string, ctx context.Context, message interface{}, params ...interface{}) {
	l.logEntryContext(info, ctx, WithPrefix(prefix, message), l.colored(`INFO`), params...)
}

func (l *prefixedLogger) DebugContext(prefix string, ctx context.Context, message interface{}, params ...interface{}) {
	l.logEntryContext(debug, ctx, WithPrefix(prefix, message), l.colored(`DEBUG`), params...)
}

func (l *prefixedLogger) TraceContext(prefix string, ctx context.Context, message interface{}, params ...interface{}) {
	l.logEntryContext(trace, ctx, WithPrefix(prefix, message), l.colored(`TRACE`), params...)
}

func (l *prefixedLogger) Error(prefix string, message interface{}, params ...interface{}) {
	l.logEntry(err, uuid.New(), WithPrefix(prefix, message), l.colored(`ERROR`), params...)
}

func (l *prefixedLogger) Warn(prefix string, message interface{}, params ...interface{}) {
	l.logEntry(warn, uuid.New(), WithPrefix(prefix, message), l.colored(`WARN`), params...)
}

func (l *prefixedLogger) Info(prefix string, message interface{}, params ...interface{}) {
	l.logEntry(info, uuid.New(), WithPrefix(prefix, message), l.colored(`INFO`), params...)
}

func (l *prefixedLogger) Debug(prefix string, message interface{}, params ...interface{}) {
	l.logEntry(debug, uuid.New(), WithPrefix(prefix, message), l.colored(`DEBUG`), params...)
}

func (l *prefixedLogger) Trace(prefix string, message interface{}, params ...interface{}) {
	l.logEntry(trace, uuid.New(), WithPrefix(prefix, message), l.colored(`TRACE`), params...)
}

func (l *prefixedLogger) Fatal(prefix string, message interface{}, params ...interface{}) {
	l.logEntry(fatal, uuid.New(), WithPrefix(prefix, message), l.colored(`FATAL`), params...)
}

func (l *prefixedLogger) Fatalln(prefix string, message interface{}, params ...interface{}) {
	l.logEntry(fatal, uuid.New(), WithPrefix(prefix, message), l.colored(`FATAL`), params...)
}

func (l *prefixedLogger) FatalContext(prefix string, ctx context.Context, message interface{}, params ...interface{}) {
	l.logEntry(fatal, uuid.New(), WithPrefix(prefix, message), l.colored(`FATAL`), params)
}

func (l *prefixedLogger)  Print(v ...interface{}){
	l.logEntry(info, uuidFromContext(context.Background()), v, l.colored(`INFO`))
}

func (l *prefixedLogger)  Printf(format string, v ...interface{}){
	l.logEntry(info, uuidFromContext(context.Background()), fmt.Sprintf(format, v), l.colored(`INFO`))
}

func (l *prefixedLogger)  Println(v ...interface{}){
	l.logEntry(info, uuidFromContext(context.Background()), v, l.colored(`INFO`))
}