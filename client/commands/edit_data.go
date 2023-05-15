package commands

import (
	"github.com/spf13/cobra"
	"gophkeeper/client/storage/model"
)

var (
	EditedDataTitle      string
	EditedMetadata       string
	EditedLocalID        string
	EditedUserCredential model.LoginPassword
	EditedUserText       model.Text
	EditedUserBinaryData model.Binary
	EditedUserBankCard   model.BankCard
)

// EditDataCmd изменяет данные.
var EditDataCmd = &cobra.Command{
	Use:   "edit",
	Short: "edit изменяет пользовательские данные по внутреннему ID",
	Long: `edit изменяет пользовательские данные по внутреннему ID. На данный момент поддерживаются 4 типа данных:
			логин пароль, текстовые данные, бинарные данные, данные банковских карт.`,
}

// EditCredentialCmd изменяет логин и пароль.
var EditCredentialCmd = &cobra.Command{
	Use:   "credential",
	Short: "credential отвечает за изменение логина и пароля",
}

// EditTextCmd изменяет пользовательские текстовые данные.
var EditTextCmd = &cobra.Command{
	Use:   "text",
	Short: "text отвечает за изменение пользовательских текстовых данных",
}

// EditBinaryCmd изменяет бинарные данные.
var EditBinaryCmd = &cobra.Command{
	Use:   "binary",
	Short: "text отвечает за изменение бинарных данных",
}

// EditBankCardCmd изменяет банковские данные.
var EditBankCardCmd = &cobra.Command{
	Use:   "bankcard",
	Short: "bankcard отвечает за изменение данных банковских карт",
}

func init() {
	EditDataCmd.PersistentFlags().StringVarP(&EditedDataTitle, "title", "t", "", "data title")
	EditDataCmd.PersistentFlags().StringVarP(&EditedMetadata, "metadata", "m", "", "metadata")
	EditDataCmd.PersistentFlags().StringVarP(&EditedLocalID, "ID", "i", "", "local ID")

	EditedUserCredential = model.LoginPassword{}
	EditCredentialCmd.Flags().StringVarP(&EditedUserCredential.Login, "login", "l", "", "data login")
	EditCredentialCmd.Flags().StringVarP(&EditedUserCredential.Password, "password", "p", "", "data password")
	EditDataCmd.AddCommand(EditCredentialCmd)

	EditedUserText = model.Text{}
	EditTextCmd.Flags().StringVarP(&EditedUserText.TextData, "text", "x", "", "data text")
	EditDataCmd.AddCommand(EditTextCmd)

	EditedUserBinaryData = model.Binary{}
	EditBinaryCmd.Flags().StringVarP(&EditedUserBinaryData.Path, "path", "p", "", "data binary path")
	EditDataCmd.AddCommand(EditBinaryCmd)

	EditedUserBankCard = model.BankCard{}
	EditBankCardCmd.Flags().StringVarP(&EditedUserBankCard.ExpDate, "date", "d", "", "bank card exp date")
	EditBankCardCmd.Flags().StringVarP(&EditedUserBankCard.Owner, "owner", "o", "", "bank card owner")
	EditBankCardCmd.Flags().StringVarP(&EditedUserBankCard.Number, "number", "n", "", "bank card number")
	EditDataCmd.AddCommand(EditBankCardCmd)
}
