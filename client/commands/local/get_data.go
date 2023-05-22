package local

import (
	"github.com/spf13/cobra"
)

var DataIndex string

// GetDataCmd возвращает выбранные данные
var GetDataCmd = &cobra.Command{
	Use:   "get",
	Short: "get выводит выбранные данные",
}

func init() {
	GetDataCmd.Flags().StringVarP(&DataIndex, "data", "d", "", "data index")
}
