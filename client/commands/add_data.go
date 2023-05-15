package commands

import (
	"github.com/spf13/cobra"
	"gophkeeper/client/storage/model"
)

var (
	DataTitle      string
	Metadata       string
	UserCredential model.LoginPassword
	UserText       model.Text
	UserBinaryData model.Binary
	UserBankCard   model.BankCard
)

// AddDataCmd добавляет данные.
var AddDataCmd = &cobra.Command{
	Use:   "add",
	Short: "Add добавляет пользовательские данные",
	Long: `Add добавляет пользовательские данные. На данный момент поддерживаются 4 типа данных:
			логин пароль, текстовые данные, бинарные данные, данные банковских карт.`,
}

// CredentialCmd добавляет логин и пароль.
var CredentialCmd = &cobra.Command{
	Use:   "credential",
	Short: "credential отвечает за добавление логина и пароля",
}

// TextCmd добавляет пользовательские текстовые данные.
var TextCmd = &cobra.Command{
	Use:   "text",
	Short: "text отвечает за добавление пользовательских текстовых данных",
}

// BinaryCmd добавляет бинарные данные.
var BinaryCmd = &cobra.Command{
	Use:   "binary",
	Short: "text отвечает за добавление бинарных данных",
}

// BankCardCmd добавляет банковские данные.
var BankCardCmd = &cobra.Command{
	Use:   "bankcard",
	Short: "bankcard отвечает за добавление данных банковских карт",
}

func init() {
	AddDataCmd.PersistentFlags().StringVarP(&DataTitle, "title", "t", "", "data title")
	AddDataCmd.PersistentFlags().StringVarP(&Metadata, "metadata", "m", "", "metadata")

	UserCredential = model.LoginPassword{}
	CredentialCmd.Flags().StringVarP(&UserCredential.Login, "login", "l", "", "data login")
	CredentialCmd.Flags().StringVarP(&UserCredential.Password, "password", "p", "", "data password")
	AddDataCmd.AddCommand(CredentialCmd)

	UserText = model.Text{}
	TextCmd.Flags().StringVarP(&UserText.TextData, "text", "x", "", "data text")
	AddDataCmd.AddCommand(TextCmd)

	UserBinaryData = model.Binary{}
	BinaryCmd.Flags().StringVarP(&UserBinaryData.Path, "path", "p", "", "data binary path")
	AddDataCmd.AddCommand(BinaryCmd)

	UserBankCard = model.BankCard{}
	BankCardCmd.Flags().StringVarP(&UserBankCard.ExpDate, "date", "d", "", "bank card exp date")
	BankCardCmd.Flags().StringVarP(&UserBankCard.Owner, "owner", "o", "", "bank card owner")
	BankCardCmd.Flags().StringVarP(&UserBankCard.Number, "number", "n", "", "bank card number")
	AddDataCmd.AddCommand(BankCardCmd)
}
