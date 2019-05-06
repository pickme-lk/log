package log

import (
	"context"
	"fmt"
	. "github.com/logrusorgru/aurora"
	"io"
	"log"
	"os"
)

type Level string

var Constructor = NewLog(FileDepth(2))

var StdLogger = Constructor.Log(FileDepth(3))

func Fatal(message interface{}, params ...interface{}) {
	StdLogger.Fatal(message, params...)
}

func Error(message interface{}, params ...interface{}) {
	StdLogger.Error(message, params...)
}

func Warn(message interface{}, params ...interface{}) {
	StdLogger.Warn(message, params...)
}

func Debug(message interface{}, params ...interface{}) {
	StdLogger.Debug(message, params...)
}

func Info(message interface{}, params ...interface{}) {
	StdLogger.Info(message, params...)
}

func Trace(message interface{}, params ...interface{}) {
	StdLogger.Trace(message, params...)
}

func FatalContext(ctx context.Context, message interface{}, params ...interface{}) {
	StdLogger.FatalContext(ctx, message, params...)
}

func ErrorContext(ctx context.Context, message interface{}, params ...interface{}) {
	StdLogger.ErrorContext(ctx, message, params...)
}

func WarnContext(ctx context.Context, message interface{}, params ...interface{}) {
	StdLogger.WarnContext(ctx, message, params...)
}

func DebugContext(ctx context.Context, message interface{}, params ...interface{}) {
	StdLogger.DebugContext(ctx, message, params...)
}

func InfoContext(ctx context.Context, message interface{}, params ...interface{}) {
	StdLogger.InfoContext(ctx, message, params...)
}

func TraceContext(ctx context.Context, message interface{}, params ...interface{}) {
	StdLogger.TraceContext(ctx, message, params...)
}

const (
	FATAL Level = `FATAL`
	ERROR Level = `ERROR`
	WARN  Level = `WARN`
	INFO  Level = `INFO`
	DEBUG Level = `DEBUG`
	TRACE Level = `TRACE`
)

var (
//fatal = `FATAL`
//err   = `ERROR`
//warn  = `WARN`
//info  = `INFO`
//debug = `DEBUG`
//trace = `TRACE`
)

var logColors = map[Level]string{
	FATAL: BgRed(`[FATAL]`).String(),
	ERROR: BgRed(`[ERROR]`).String(),
	WARN:  BgBrown(`[WARN]`).String(),
	INFO:  BgBlue(`[INFO]`).String(),
	DEBUG: BgCyan(`[DEBUG]`).String(),
	TRACE: BgMagenta(`[TRACE]`).String(),
}

var logTypes = map[Level]int{
	FATAL: 0,
	ERROR: 1,
	WARN:  2,
	INFO:  3,
	DEBUG: 4,
	TRACE: 5,
}

type SimpleLogger interface {
	Print(v ...interface{})
	Printf(format string, v ...interface{})
	Println(v ...interface{})
}

type Logger interface {
	Fatal(message interface{}, params ...interface{})
	Error(message interface{}, params ...interface{})
	Warn(message interface{}, params ...interface{})
	Debug(message interface{}, params ...interface{})
	Info(message interface{}, params ...interface{})
	Trace(message interface{}, params ...interface{})
	FatalContext(ctx context.Context, message interface{}, params ...interface{})
	ErrorContext(ctx context.Context, message interface{}, params ...interface{})
	WarnContext(ctx context.Context, message interface{}, params ...interface{})
	DebugContext(ctx context.Context, message interface{}, params ...interface{})
	InfoContext(ctx context.Context, message interface{}, params ...interface{})
	TraceContext(ctx context.Context, message interface{}, params ...interface{})
	SimpleLogger
}

type PrefixedLogger interface {
	Fatal(prefix string, message interface{}, params ...interface{})
	Error(prefix string, message interface{}, params ...interface{})
	Warn(prefix string, message interface{}, params ...interface{})
	Debug(prefix string, message interface{}, params ...interface{})
	Info(prefix string, message interface{}, params ...interface{})
	Trace(prefix string, message interface{}, params ...interface{})
	FatalContext(ctx context.Context, prefix string, message interface{}, params ...interface{})
	ErrorContext(ctx context.Context, prefix string, message interface{}, params ...interface{})
	WarnContext(ctx context.Context, prefix string, message interface{}, params ...interface{})
	DebugContext(ctx context.Context, prefix string, message interface{}, params ...interface{})
	InfoContext(ctx context.Context, prefix string, message interface{}, params ...interface{})
	TraceContext(ctx context.Context, prefix string, message interface{}, params ...interface{})
	SimpleLogger
}

type Log interface {
	Log(...Option) Logger
	SimpleLog() SimpleLogger
	PrefixedLog(...Option) PrefixedLogger
}

type logIpl struct {
	log *log.Logger
	*logOptions
}

func NewLog(options ...Option) Log {
	opts := new(logOptions)
	opts.applyDefault()
	opts.apply(options...)

	return &logIpl{
		log:        log.New(opts.writer, ``, log.LstdFlags|log.Lmicroseconds),
		logOptions: opts,
	}
}

type logOptions struct {
	prefix    string
	colors    bool
	logLevel  Level
	filePath  bool
	fileDepth int
	writer    io.Writer
}

func (lOpts *logOptions) applyDefault() {
	lOpts.fileDepth = 2
	lOpts.colors = true
	lOpts.logLevel = TRACE
	lOpts.filePath = true
	lOpts.writer = os.Stdout
}

func (lOpts *logOptions) copy() *logOptions {
	return &logOptions{
		fileDepth: lOpts.fileDepth,
		colors:    lOpts.colors,
		logLevel:  lOpts.logLevel,
		filePath:  lOpts.filePath,
		writer:    lOpts.writer,
	}
}

func (lOpts *logOptions) apply(options ...Option) {
	for _, opt := range options {
		opt(lOpts)
	}
}

type Option func(*logOptions)

func FileDepth(d int) Option {
	return func(opts *logOptions) {
		opts.fileDepth = d
	}
}

func WithStdOut(w io.Writer) Option {
	return func(opts *logOptions) {
		opts.writer = w
	}
}

func WithFilePath(enabled bool) Option {
	return func(opts *logOptions) {
		opts.filePath = enabled
	}
}

func Prefixed(prefix string) Option {
	return func(opts *logOptions) {
		opts.prefix = prefix + `.`
	}
}

func WithColors(enabled bool) Option {
	return func(opts *logOptions) {
		opts.colors = enabled
	}
}

func WithLevel(level Level) Option {
	return func(opts *logOptions) {
		opts.logLevel = level
	}
}

func (l *logIpl) Log(options ...Option) Logger {

	opts := l.logOptions.copy()
	opts.apply(options...)

	return &logger{
		logParser: logParser{
			logOptions: opts,
			log:        l.log,
		},
	}
}

func (*logIpl) SimpleLog() SimpleLogger {
	panic(`implement me`)
}

func (l *logIpl) PrefixedLog(options ...Option) PrefixedLogger {
	opts := l.logOptions.copy()
	opts.apply(options...)

	return &prefixedLogger{
		logParser: logParser{
			logOptions: opts,
			log:        l.log,
		},
	}
}

type logger struct {
	logParser
}

func (l *logger) ErrorContext(ctx context.Context, message interface{}, params ...interface{}) {
	l.logEntry(ERROR, ctx, message, params...)
}

func (l *logger) WarnContext(ctx context.Context, message interface{}, params ...interface{}) {
	l.logEntry(WARN, ctx, message, params...)
}

func (l *logger) InfoContext(ctx context.Context, message interface{}, params ...interface{}) {
	l.logEntry(INFO, ctx, message, params...)
}

func (l *logger) DebugContext(ctx context.Context, message interface{}, params ...interface{}) {
	l.logEntry(DEBUG, ctx, message, params...)
}

func (l *logger) TraceContext(ctx context.Context, message interface{}, params ...interface{}) {
	l.logEntry(TRACE, ctx, message, params...)
}

func (l *logger) Error(message interface{}, params ...interface{}) {
	l.logEntry(ERROR, nil, message, params...)
}

func (l *logger) Warn(message interface{}, params ...interface{}) {
	l.logEntry(WARN, nil, message, params...)
}

func (l *logger) Info(message interface{}, params ...interface{}) {
	l.logEntry(INFO, nil, message, params...)
}

func (l *logger) Debug(message interface{}, params ...interface{}) {
	l.logEntry(DEBUG, nil, message, params...)
}

func (l *logger) Trace(message interface{}, params ...interface{}) {
	l.logEntry(TRACE, nil, message, params...)
}

func (l *logger) Fatal(message interface{}, params ...interface{}) {
	l.logEntry(FATAL, nil, message, params...)
}

func (l *logger) Fatalln(message interface{}, params ...interface{}) {
	l.logEntry(FATAL, nil, message, params...)
}

func (l *logger) FatalContext(ctx context.Context, message interface{}, params ...interface{}) {
	l.logEntry(FATAL, nil, message, params)
}

func (l *logger) Print(v ...interface{}) {
	l.logEntry(INFO, nil, v, `INFO`)
}

func (l *logger) Printf(format string, v ...interface{}) {
	l.logEntry(INFO, nil, fmt.Sprintf(format, v...), `INFO`)
}

func (l *logger) Println(v ...interface{}) {
	l.logEntry(INFO, nil, v, `INFO`)
}
