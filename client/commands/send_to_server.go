package commands

import "github.com/spf13/cobra"

// SendToServerCmd represents the login command
var SendToServerCmd = &cobra.Command{
	Use:   "send",
	Short: "send отправляет данные на сервер",
}

func init() {
	SendToServerCmd.Flags().StringVarP(&Masterkey, "masterkey", "m", "", "master key")
}
