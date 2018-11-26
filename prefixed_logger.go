package log

import (
	"context"
	"fmt"
	"github.com/google/uuid"
)

type prefixedLogger struct {
	logParser
}

func (l *prefixedLogger) ErrorContext(ctx context.Context, prefix string, message interface{}, params ...interface{}) {
	l.logEntry(err, uuidFromContext(ctx), WithPrefix(prefix, message), l.colored(`ERROR`), params...)
}

func (l *prefixedLogger) WarnContext(ctx context.Context, prefix string, message interface{}, params ...interface{}) {
	l.logEntry(warn, uuidFromContext(ctx), WithPrefix(prefix, message), l.colored(`WARN`), params...)
}

func (l *prefixedLogger) InfoContext(ctx context.Context, prefix string, message interface{}, params ...interface{}) {
	l.logEntry(info, uuidFromContext(ctx), WithPrefix(prefix, message), l.colored(`INFO`), params...)
}

func (l *prefixedLogger) DebugContext(ctx context.Context, prefix string, message interface{}, params ...interface{}) {
	l.logEntry(debug, uuidFromContext(ctx), WithPrefix(prefix, message), l.colored(`DEBUG`), params...)
}

func (l *prefixedLogger) TraceContext(ctx context.Context, prefix string, message interface{}, params ...interface{}) {
	l.logEntry(trace, uuidFromContext(ctx), WithPrefix(prefix, message), l.colored(`TRACE`), params...)
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

func (l *prefixedLogger) FatalContext(ctx context.Context, prefix string, message interface{}, params ...interface{}) {
	l.logEntry(fatal, uuid.New(), WithPrefix(prefix, message), l.colored(`FATAL`), params)
}

func (l *prefixedLogger) Print(v ...interface{}) {
	l.logEntry(info, uuidFromContext(context.Background()), v, l.colored(`INFO`))
}

func (l *prefixedLogger) Printf(format string, v ...interface{}) {
	l.logEntry(info, uuidFromContext(context.Background()), fmt.Sprintf(format, v), l.colored(`INFO`))
}

func (l *prefixedLogger) Println(v ...interface{}) {
	l.logEntry(info, uuidFromContext(context.Background()), v, l.colored(`INFO`))
}
