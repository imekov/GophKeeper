package commands

import (
	"github.com/spf13/cobra"
)

// LoginCmd represents the login command
var LoginCmd = &cobra.Command{
	Use:   "login",
	Short: "Авторизация в сервисе GophKeeper",
	Long: `Авторизация в сервисе GophKeeper. Для авторизации нужно ввести
			пару логин и пароль`,
}

func init() {
	LoginCmd.Flags().StringVarP(&Login, "login", "l", "", "user login")
	LoginCmd.Flags().StringVarP(&Password, "password", "p", "", "user password")
}
