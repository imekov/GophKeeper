package handlers

import (
	"context"
	"fmt"

	pb "gophkeeper/proto"

	"github.com/spf13/cobra"
)

// Register используется для регистрации нового пользователя.
func (h Handlers) Register(login *string, password *string) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		newUser := pb.AuthRequest{Login: *login, Password: *password}
		resp, err := h.Client.Register(context.Background(), &newUser)
		if err != nil {
			fmt.Printf("Возникла ошибка при регистрации: %s", err.Error())
			return
		}
		h.Repo.UpdateToken(resp.Token)
		fmt.Println("Регистрация прошла успешно")

	}
}

// Login используется для авторизации пользователей.
func (h Handlers) Login(login *string, password *string) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		userCredentials := pb.AuthRequest{Login: *login, Password: *password}
		resp, err := h.Client.Login(context.Background(), &userCredentials)
		if err != nil {
			fmt.Printf("Возникла ошибка при авторизации: %s", err.Error())
			return
		}
		h.Repo.UpdateToken(resp.Token)
		fmt.Println("Авторизация прошла успешно")
	}

}
