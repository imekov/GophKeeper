package server

import (
	"github.com/spf13/cobra"

	"gophkeeper/client/commands"
)

// LoginCmd производит авторизацию
var LoginCmd = &cobra.Command{
	Use:   "login",
	Short: "Авторизация в сервисе GophKeeper",
	Long: `Авторизация в сервисе GophKeeper. Для авторизации нужно ввести
			пару логин и пароль`,
}

func init() {
	LoginCmd.Flags().StringVarP(&commands.Login, "login", "l", "", "user login")
	LoginCmd.Flags().StringVarP(&commands.Password, "password", "p", "", "user password")
}
