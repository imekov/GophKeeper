package main

import (
	"log"

	_ "github.com/lib/pq"

	"gophkeeper/server"
	"gophkeeper/server/config"
	"gophkeeper/server/logger"
)

func main() {
	cfg := config.GetConfig()
	mainLogger := logger.NewLogger(cfg.LoggerDebug)

	connection := server.DBConnectionInitialization(cfg.DatabaseURI, &mainLogger)
	defer connection.DBCloseConnection()

	if err := connection.StartGRPCServer(cfg.RunAddress, &mainLogger); err != nil {
		log.Fatal(err)
	}

}
