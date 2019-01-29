package log

import (
	"context"
	"fmt"
)

type prefixedLogger struct {
	logParser
}

func (l *prefixedLogger) ErrorContext(ctx context.Context, prefix string, message interface{}, params ...interface{}) {
	l.logEntry(ERROR, ctx, WithPrefix(prefix, message), params...)
}

func (l *prefixedLogger) WarnContext(ctx context.Context, prefix string, message interface{}, params ...interface{}) {
	l.logEntry(WARN, ctx, WithPrefix(prefix, message), params...)
}

func (l *prefixedLogger) InfoContext(ctx context.Context, prefix string, message interface{}, params ...interface{}) {
	l.logEntry(INFO, ctx, WithPrefix(prefix, message), params...)
}

func (l *prefixedLogger) DebugContext(ctx context.Context, prefix string, message interface{}, params ...interface{}) {
	l.logEntry(DEBUG, ctx, WithPrefix(prefix, message), params...)
}

func (l *prefixedLogger) TraceContext(ctx context.Context, prefix string, message interface{}, params ...interface{}) {
	l.logEntry(TRACE, ctx, WithPrefix(prefix, message), params...)
}

func (l *prefixedLogger) Error(prefix string, message interface{}, params ...interface{}) {
	l.logEntry(ERROR, nil, WithPrefix(prefix, message), params...)
}

func (l *prefixedLogger) Warn(prefix string, message interface{}, params ...interface{}) {
	l.logEntry(WARN, nil, WithPrefix(prefix, message), params...)
}

func (l *prefixedLogger) Info(prefix string, message interface{}, params ...interface{}) {
	l.logEntry(INFO, nil, WithPrefix(prefix, message), params...)
}

func (l *prefixedLogger) Debug(prefix string, message interface{}, params ...interface{}) {
	l.logEntry(DEBUG, nil, WithPrefix(prefix, message), params...)
}

func (l *prefixedLogger) Trace(prefix string, message interface{}, params ...interface{}) {
	l.logEntry(TRACE, nil, WithPrefix(prefix, message), params...)
}

func (l *prefixedLogger) Fatal(prefix string, message interface{}, params ...interface{}) {
	l.logEntry(FATAL, nil, WithPrefix(prefix, message), params...)
}

func (l *prefixedLogger) Fatalln(prefix string, message interface{}, params ...interface{}) {
	l.logEntry(FATAL, nil, WithPrefix(prefix, message), params...)
}

func (l *prefixedLogger) FatalContext(ctx context.Context, prefix string, message interface{}, params ...interface{}) {
	l.logEntry(FATAL, nil, WithPrefix(prefix, message), params)
}

func (l *prefixedLogger) Print(v ...interface{}) {
	l.logEntry(INFO, nil, v, l.colored(`INFO`))
}

func (l *prefixedLogger) Printf(format string, v ...interface{}) {
	l.logEntry(INFO, nil, fmt.Sprintf(format, v...), l.colored(`INFO`))
}

func (l *prefixedLogger) Println(v ...interface{}) {
	l.logEntry(INFO, nil, v, l.colored(`INFO`))

}
