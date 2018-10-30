package main

import (
	"gitlab.mytaxi.lk/pickme/go/log"
	log2 "gitlab.mytaxi.lk/pickme/go/log/example/log"
)

func main() {

	log.Constructor = log.NewLog(log.WithLevel(log.INFO))

	log.StdLogger.Info(`sadasdasd`)

	log.Info(`sadasdasd`)

	otherLogger := log2.NewOtherLogger()

	otherLogger.Info(123123)
	otherLogger.Trace(123123)

}
