package client

import (
	"log"

	"gophkeeper/client/commands"
	"gophkeeper/client/commands/local"
	"gophkeeper/client/commands/server"
	"gophkeeper/client/handlers"
	"gophkeeper/client/storage/filesystem"
	pb "gophkeeper/proto"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Client содержит ссылку на рут команду, хэндлеры и grpc клиент.
type Client struct {
	RootCmd     *cobra.Command
	handlers    handlers.Handlers
	GRPCConnect *grpc.ClientConn
}

// NewClient создает новый экземпляр клиента.
func NewClient(serverAddress string, filename string) *Client {

	conn, err := grpc.Dial(serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	c := pb.NewGophKeeperClient(conn)
	f := filesystem.NewFSStorage(filename)

	newHandler := handlers.Handlers{
		Client: c,
		Repo:   f,
	}

	newClient := Client{
		RootCmd:     commands.RootCmd,
		handlers:    newHandler,
		GRPCConnect: conn,
	}

	local.GetListCmd.Run = newClient.handlers.GetUserDataList()
	newClient.RootCmd.AddCommand(local.GetListCmd)

	local.GetDataCmd.Run = newClient.handlers.GetUserData(&local.DataIndex)
	newClient.RootCmd.AddCommand(local.GetDataCmd)

	local.CredentialCmd.Run = newClient.handlers.AddData(&local.DataTitle, &local.Metadata, &local.UserCredential)
	local.TextCmd.Run = newClient.handlers.AddData(&local.DataTitle, &local.Metadata, &local.UserText)
	local.BankCardCmd.Run = newClient.handlers.AddData(&local.DataTitle, &local.Metadata, &local.UserBankCard)
	local.BinaryCmd.PreRun = newClient.handlers.MakeBinary(&local.UserBinaryData)
	local.BinaryCmd.Run = newClient.handlers.AddData(&local.DataTitle, &local.Metadata, &local.UserBinaryData)
	newClient.RootCmd.AddCommand(local.AddDataCmd)

	local.EditCredentialCmd.Run = newClient.handlers.EditData(&local.EditedLocalID, &local.EditedDataTitle, &local.EditedMetadata, &local.EditedUserCredential)
	local.EditTextCmd.Run = newClient.handlers.EditData(&local.EditedLocalID, &local.EditedDataTitle, &local.EditedMetadata, &local.EditedUserText)
	local.EditBankCardCmd.Run = newClient.handlers.EditData(&local.EditedLocalID, &local.EditedDataTitle, &local.EditedMetadata, &local.EditedUserBankCard)
	local.EditBinaryCmd.PreRun = newClient.handlers.MakeBinary(&local.EditedUserBinaryData)
	local.EditBinaryCmd.Run = newClient.handlers.EditData(&local.EditedLocalID, &local.EditedDataTitle, &local.EditedMetadata, &local.EditedUserBinaryData)
	newClient.RootCmd.AddCommand(local.EditDataCmd)

	server.RegisterCmd.Run = newClient.handlers.Register(&commands.Login, &commands.Password)
	newClient.RootCmd.AddCommand(server.RegisterCmd)

	server.LoginCmd.Run = newClient.handlers.Login(&commands.Login, &commands.Password)
	newClient.RootCmd.AddCommand(server.LoginCmd)

	server.SynchronizationCmd.Run = newClient.handlers.Sync(&commands.Masterkey)
	newClient.RootCmd.AddCommand(server.SynchronizationCmd)

	return &newClient
}

// CloseConnection закрывает клиентское соединение с GRPC сервером.
func (c Client) CloseConnection() {
	c.GRPCConnect.Close()
}
