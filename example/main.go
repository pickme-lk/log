package main

import (
	"context"
	"github.com/google/uuid"
	"gitlab.mytaxi.lk/pickme/go-util/traceable_context"
	"gitlab.mytaxi.lk/pickme/go/log"
	customLog "gitlab.mytaxi.lk/pickme/go/log/example/log"
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

	log.ErrorContext(ctx, `message`, `param1`, `param2`)
	log.ErrorContext(ctx, `message`)
	log.ErrorContext(ctx, `message`)
	log.ErrorContext(ctx, log.WithPrefix(`prefix`, `message`))
	log.WarnContext(ctx, log.WithPrefix(`prefix`, `message`), `param1`, `param2`)
	// prefixed log
	prefixedLogger := log.Constructor.PrefixedLog(log.WithLevel(log.ERROR))
	prefixedLogger.Info(`module.Func`, `message`)
	prefixedLogger.Trace(`module.Func`, `message`)
	prefixedLogger.Error(`module.Func`, `message`)
	prefixedLogger.Error(`module.Func`, `message`, `param1`, `param2`)
	//
	//// custom logger
	logger := customLog.NewOtherLogger()
	logger.Info(`info`)
	logger.Trace(`trace`)

}
