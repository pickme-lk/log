package log

import (
	"context"
)

type noopLogger struct{}

func NewNoopLogger() Logger {
	return new(noopLogger)
}
func (l *noopLogger) ErrorContext(ctx context.Context, message interface{}, params ...interface{}) {}
func (l *noopLogger) WarnContext(ctx context.Context, message interface{}, params ...interface{})  {}
func (l *noopLogger) InfoContext(ctx context.Context, message interface{}, params ...interface{})  {}
func (l *noopLogger) DebugContext(ctx context.Context, message interface{}, params ...interface{}) {}
func (l *noopLogger) TraceContext(ctx context.Context, message interface{}, params ...interface{}) {}
func (l *noopLogger) Error(message interface{}, params ...interface{})                             {}
func (l *noopLogger) Warn(message interface{}, params ...interface{})                              {}
func (l *noopLogger) Info(message interface{}, params ...interface{})                              {}
func (l *noopLogger) Debug(message interface{}, params ...interface{})                             {}
func (l *noopLogger) Trace(message interface{}, params ...interface{})                             {}
func (l *noopLogger) Fatal(message interface{}, params ...interface{})                             {}
func (l *noopLogger) Fatalln(message interface{}, params ...interface{})                           {}
func (l *noopLogger) FatalContext(ctx context.Context, message interface{}, params ...interface{}) {}
func (l *noopLogger) Print(v ...interface{})                                                       {}
func (l *noopLogger) Printf(format string, v ...interface{})                                       {}
func (l *noopLogger) Println(v ...interface{})                                                     {}
func (l *noopLogger) NewLog(opts ...Option) Logger                                                 { return NewNoopLogger() }
func (l *noopLogger) NewPrefixedLog(opts ...Option) PrefixedLogger                                 { return NewPrefixedNoopLogger() }

type prefixedNoopLogger struct{}

func NewPrefixedNoopLogger() PrefixedLogger {
	return new(prefixedNoopLogger)
}
func (l *prefixedNoopLogger) ErrorContext(ctx context.Context, prefix string, message interface{}, params ...interface{}) {
}
func (l *prefixedNoopLogger) WarnContext(ctx context.Context, prefix string, message interface{}, params ...interface{}) {
}
func (l *prefixedNoopLogger) InfoContext(ctx context.Context, prefix string, message interface{}, params ...interface{}) {
}
func (l *prefixedNoopLogger) DebugContext(ctx context.Context, prefix string, message interface{}, params ...interface{}) {
}
func (l *prefixedNoopLogger) TraceContext(ctx context.Context, prefix string, message interface{}, params ...interface{}) {
}
func (l *prefixedNoopLogger) Error(prefix string, message interface{}, params ...interface{})   {}
func (l *prefixedNoopLogger) Warn(prefix string, message interface{}, params ...interface{})    {}
func (l *prefixedNoopLogger) Info(prefix string, message interface{}, params ...interface{})    {}
func (l *prefixedNoopLogger) Debug(prefix string, message interface{}, params ...interface{})   {}
func (l *prefixedNoopLogger) Trace(prefix string, message interface{}, params ...interface{})   {}
func (l *prefixedNoopLogger) Fatal(prefix string, message interface{}, params ...interface{})   {}
func (l *prefixedNoopLogger) Fatalln(prefix string, message interface{}, params ...interface{}) {}
func (l *prefixedNoopLogger) FatalContext(ctx context.Context, prefix string, message interface{}, params ...interface{}) {
}
func (l *prefixedNoopLogger) Print(v ...interface{})                 {}
func (l *prefixedNoopLogger) Printf(format string, v ...interface{}) {}
func (l *prefixedNoopLogger) Println(v ...interface{})               {}

func (l *prefixedNoopLogger) NewLog(opts ...Option) Logger { return NewNoopLogger() }
func (l *prefixedNoopLogger) NewPrefixedLog(opts ...Option) PrefixedLogger {
	return NewPrefixedNoopLogger()
}
