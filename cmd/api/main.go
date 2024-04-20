package main

import (
	"bookingtogo/config"
	"bookingtogo/internal/server"
	"bookingtogo/pkg/infra/db"
	"bookingtogo/pkg/infra/logger"
)

func main() {

	//* ====================== Config ======================

	conf := config.InitConfig("dev")

	//* ====================== Logger ======================

	//* Loggrus
	appLogger := logger.NewLogrusLogger(&conf.Logger.Logrus)

	//* ====================== Connection DB ======================

	var dbList db.DatabaseList
	dbList.DatabaseApp = db.NewGORMConnection(&conf.Connection.DatabaseApp, appLogger)

	//* ====================== Running Server ======================

	server.Run(conf, &dbList, appLogger)
}
