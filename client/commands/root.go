package commands

import (
	"github.com/spf13/cobra"
)

var Login, Password, Masterkey string

// RootCmd представляет базовую команду при вызове без каких-либо подкоманд
var RootCmd = &cobra.Command{
	Use:   "client",
	Short: "Gophkeeper хранит пользовательские данные",
	Long: `Gophkeeper хранит пользовательские данные. Для авторизации используйте
			ключ login, для регистрации register`,
}
