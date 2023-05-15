package main

import (
	"gophkeeper/server"
	"gophkeeper/server/config"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	cfg := config.GetConfig()

	connection := server.DBConnectionInitialization(cfg.DatabaseURI)
	defer connection.DBCloseConnection()

	if err := connection.StartGRPCServer(cfg.RunAddress); err != nil {
		log.Fatal(err)
	}

}
