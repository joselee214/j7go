package components

import (
	"github.com/joselee214/j7f/components/log"
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
