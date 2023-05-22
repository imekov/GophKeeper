package server

import (
	"github.com/spf13/cobra"

	"gophkeeper/client/commands"
)

// RegisterCmd реализуют регистрацию в сервисе
var RegisterCmd = &cobra.Command{
	Use:   "register",
	Short: "Регистрация в сервисе GophKeeper",
	Long: `Регистрация в сервисе GophKeeper. Для регистрации нужно ввести
			пару логин и пароль`,
}

func init() {
	RegisterCmd.Flags().StringVarP(&commands.Login, "login", "l", "", "user login")
	RegisterCmd.Flags().StringVarP(&commands.Password, "password", "p", "", "user password")
}
