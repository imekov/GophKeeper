package main

import (
	"fmt"
	"os"

	"gophkeeper/client"
	"gophkeeper/client/config"
)

func main() {
	cfg := config.GetConfig()

	newClient := client.NewClient(cfg.ServerAddress, cfg.Filename)
	defer newClient.CloseConnection()

	if err := newClient.RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
