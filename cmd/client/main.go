package main

import (
	"fmt"
	"gophkeeper/client"
	"os"
)

func main() {
	//TODO: брать данные из конфига
	newClient := client.NewClient(":8080", "client/userdata/data.gob")
	defer newClient.CloseConnection()

	if err := newClient.RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
