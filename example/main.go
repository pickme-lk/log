package main

import (
	"context"
	"github.com/google/uuid"
	"gitlab.mytaxi.lk/pickme/go-util/traceable_context"
	"gitlab.mytaxi.lk/pickme/go/log"
)

func main() {

	// usage of log
	//log.Info(`info`)
	//log.Trace(`trace`)
	//log.Error(`error`)
	//log.Error(log.WithPrefix(`prefix`, `error`))

	// log with a traceable context
	tCtx := traceable_context.WithUUID(uuid.New())

	ctx, _ := context.WithCancel(tCtx)

	log.ErrorContext(ctx, `info`)
	log.ErrorContext(ctx, `trace`)
	log.ErrorContext(ctx, `error`)
	log.ErrorContext(ctx, log.WithPrefix(`prefix`, `error`))
	// prefixed log
	//prefixedLogger := log.Constructor.PrefixedLog(log.WithLevel(log.ERROR))
	//prefixedLogger.Info(`module.Func`, `info`)
	//prefixedLogger.Trace(`module.Func`, `trace`)
	//prefixedLogger.Error(`module.Func`, `error`)
	//prefixedLogger.Error(`module.Func`, `error`)
	//
	//// custom logger
	//logger := customLog.NewOtherLogger()
	//logger.Info(`info`)
	//logger.Trace(`trace`)

}
