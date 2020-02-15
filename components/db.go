package components

import (
	"github.com/joselee214/j7f/components/dao"
	"go.uber.org/zap"
)

var M *dao.Node

func InitDB(cfg *dao.DBConfig) error {
	var err error
	M, err = dao.NewNode(cfg, checkErr)
	if err != nil {
		return err
	}
	return nil
}

func checkErr(err error) {
	L.Error("DB err", zap.String("db", err.Error()))
}
