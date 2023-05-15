package commands

import "github.com/spf13/cobra"

// GetFromServerCmd принимает данные с сервера.
var GetFromServerCmd = &cobra.Command{
	Use:   "download",
	Short: "download принимает данные с сервера",
}

func init() {
	GetFromServerCmd.Flags().StringVarP(&Masterkey, "masterkey", "m", "", "master key")
}
