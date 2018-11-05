package main

import (
	"github.com/google/uuid"
	"gitlab.mytaxi.lk/pickme/go-util/traceable_context"
	"gitlab.mytaxi.lk/pickme/go/log"
	customLog "gitlab.mytaxi.lk/pickme/go/log/example/log"
)

func main() {

	// usage of log
	log.Info(`info`)
	log.Trace(`trace`)
	log.Error(`error`)
	log.Error(log.WithPrefix(`prefix`, `error`))

	// log with a traceable context
	ctx := traceable_context.WithUUID(uuid.New())
	log.ErrorContext(ctx, `info`)
	log.ErrorContext(ctx, `trace`)
	log.ErrorContext(ctx, `error`)
	log.ErrorContext(ctx, log.WithPrefix(`prefix`, `error`))
	// prefixed log
	prefixedLogger := log.Constructor.PrefixedLog(log.WithLevel(log.ERROR))
	prefixedLogger.Info(`module.Func`, `info`)
	prefixedLogger.Trace(`module.Func`, `trace`)
	prefixedLogger.Error(`module.Func`, `error`)
	prefixedLogger.Error(`module.Func`, `error`)

	// custom logger
	logger := customLog.NewOtherLogger()
	logger.Info(`info`)
	logger.Trace(`trace`)

}
