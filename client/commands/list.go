package commands

import (
	"github.com/spf13/cobra"
)

// GetListCmd возвращает список сохраненных данных
var GetListCmd = &cobra.Command{
	Use:   "list",
	Short: "list отображает список данных",
}
