package components

import (
	"go.7yes.com/go/components/log"
)

var L *log.Logger

func InitLog(logConfig *log.Config) (err error) {
	logger, err := log.NewZap(logConfig)
	L = logger
	return err
}

func ResetLog(logConfig *log.Config) (err error) {
	logger, err := log.NewZap(logConfig)
	L.ResetLogger(logger)
	return err
}
