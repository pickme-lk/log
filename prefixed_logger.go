package log

import (
	"context"
	"fmt"
)

type prefixedLogger struct {
	logParser
}

func (l *prefixedLogger) ErrorContext(ctx context.Context, prefix string, message interface{}, params ...interface{}) {
	l.logEntry(ERROR, ctx, l.WithPrefix(prefix, message), params...)
}

func (l *prefixedLogger) WarnContext(ctx context.Context, prefix string, message interface{}, params ...interface{}) {
	l.logEntry(WARN, ctx, l.WithPrefix(prefix, message), params...)
}

func (l *prefixedLogger) InfoContext(ctx context.Context, prefix string, message interface{}, params ...interface{}) {
	l.logEntry(INFO, ctx, l.WithPrefix(prefix, message), params...)
}

func (l *prefixedLogger) DebugContext(ctx context.Context, prefix string, message interface{}, params ...interface{}) {
	l.logEntry(DEBUG, ctx, l.WithPrefix(prefix, message), params...)
}

func (l *prefixedLogger) TraceContext(ctx context.Context, prefix string, message interface{}, params ...interface{}) {
	l.logEntry(TRACE, ctx, l.WithPrefix(prefix, message), params...)
}

func (l *prefixedLogger) Error(prefix string, message interface{}, params ...interface{}) {
	l.logEntry(ERROR, nil, l.WithPrefix(prefix, message), params...)
}

func (l *prefixedLogger) Warn(prefix string, message interface{}, params ...interface{}) {
	l.logEntry(WARN, nil, l.WithPrefix(prefix, message), params...)
}

func (l *prefixedLogger) Info(prefix string, message interface{}, params ...interface{}) {
	l.logEntry(INFO, nil, l.WithPrefix(prefix, message), params...)
}

func (l *prefixedLogger) Debug(prefix string, message interface{}, params ...interface{}) {
	l.logEntry(DEBUG, nil, l.WithPrefix(prefix, message), params...)
}

func (l *prefixedLogger) Trace(prefix string, message interface{}, params ...interface{}) {
	l.logEntry(TRACE, nil, l.WithPrefix(prefix, message), params...)
}

func (l *prefixedLogger) Fatal(prefix string, message interface{}, params ...interface{}) {
	l.logEntry(FATAL, nil, l.WithPrefix(prefix, message), params...)
}

func (l *prefixedLogger) Fatalln(prefix string, message interface{}, params ...interface{}) {
	l.logEntry(FATAL, nil, l.WithPrefix(prefix, message), params...)
}

func (l *prefixedLogger) FatalContext(ctx context.Context, prefix string, message interface{}, params ...interface{}) {
	l.logEntry(FATAL, nil, l.WithPrefix(prefix, message), params)
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

func (l *prefixedLogger) NewLog(opts ...Option) Logger {
	defaults := l.logOptions.copy()
	defaults.apply(opts...)

	return &logger{
		logParser: logParser{
			logOptions: defaults,
			log:        l.log,
		},
	}
}

func (l *prefixedLogger) NewPrefixedLog(opts ...Option) PrefixedLogger {
	defaults := l.logOptions.copy()
	defaults.apply(opts...)

	return &prefixedLogger{
		logParser: logParser{
			logOptions: defaults,
			log:        l.log,
		},
	}
}
