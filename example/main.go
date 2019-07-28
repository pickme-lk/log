package main

import (
	"context"
	"github.com/google/uuid"
	"github.com/pickme-go/log"
	customLog "github.com/pickme-go/log/example/log"
	"github.com/pickme-go/traceable-context"
)

func main() {
	// usage of log
	log.Info(`message`, `param1`, `param2`)
	log.Trace(`message`)
	log.Error(`message`)
	log.Error(log.WithPrefix(`prefix`, `message`), `param1`, `param2`)

	// log with a traceable context
	tCtx := traceable_context.WithUUID(uuid.New())
	ctx, _ := context.WithCancel(tCtx)
	logger := log.Constructor.Log(log.WithColors(true), log.WithLevel(log.TRACE), log.WithFilePath(false), log.Prefixed(`level-1`))
	logger.ErrorContext(ctx, `message`, `param1`, `param2`)
	logger.ErrorContext(ctx, `message`)
	logger.ErrorContext(ctx, `message`)
	logger.ErrorContext(ctx, log.WithPrefix(`prefix`, `message`))
	logger.WarnContext(ctx, log.WithPrefix(`prefix`, `message`), `param1`, `param2`)

	// prefixed log`
	prefixedLogger := log.Constructor.PrefixedLog(log.WithLevel(log.ERROR), log.WithFilePath(false))
	prefixedLogger.Info(`module.sub-module`, `message`)
	prefixedLogger.Trace(`module.sub-module`, `message`)
	prefixedLogger.Error(`module.sub-module`, `message`)
	prefixedLogger.Error(`module.sub-module`, `message`, `param1`, `param2`)

	// custom logger
	customLogger := customLog.NewLogger()
	customLogger.Info(`info`)
	customLogger.Trace(`trace`)

	// create a logger instance derived from logger
	nestedLogger := logger.NewLog(log.WithLevel(log.TRACE), log.Prefixed(`level-2`))
	nestedLogger.Error(`error happened`, 22)

}
