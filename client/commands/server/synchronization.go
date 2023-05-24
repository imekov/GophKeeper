package server

import (
	"github.com/spf13/cobra"

	"gophkeeper/client/commands"
)

// SynchronizationCmd синхронизирует данные
var SynchronizationCmd = &cobra.Command{
	Use:   "sync",
	Short: "sync синхронизирует данные между клиентом и сервером",
	Long: `sync синхронизирует данные между клиентом и сервером
			Для работы нужно указать параметр -m с ключом для шифровки/расшифровки данных`,
}

func init() {
	SynchronizationCmd.Flags().StringVarP(&commands.Masterkey, "masterkey", "m", "", "master key")
}
