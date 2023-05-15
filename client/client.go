package client

import (
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gophkeeper/client/commands"
	"gophkeeper/client/handlers"
	"gophkeeper/client/storage/filesystem"
	pb "gophkeeper/proto"
	"log"
)

type Client struct {
	RootCmd     *cobra.Command
	handlers    handlers.Handlers
	GRPCConnect *grpc.ClientConn
}

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

	commands.RegisterCmd.Run = newClient.handlers.Register(&commands.Login, &commands.Password)
	newClient.RootCmd.AddCommand(commands.RegisterCmd)

	commands.LoginCmd.Run = newClient.handlers.Login(&commands.Login, &commands.Password)
	newClient.RootCmd.AddCommand(commands.LoginCmd)

	commands.GetListCmd.Run = newClient.handlers.GetUserDataList()
	newClient.RootCmd.AddCommand(commands.GetListCmd)

	commands.GetDataCmd.Run = newClient.handlers.GetUserData(&commands.DataIndex)
	newClient.RootCmd.AddCommand(commands.GetDataCmd)

	commands.CredentialCmd.Run = newClient.handlers.AddData(&commands.DataTitle, &commands.Metadata, &commands.UserCredential)
	commands.TextCmd.Run = newClient.handlers.AddData(&commands.DataTitle, &commands.Metadata, &commands.UserText)
	commands.BankCardCmd.Run = newClient.handlers.AddData(&commands.DataTitle, &commands.Metadata, &commands.UserBankCard)
	commands.BinaryCmd.PreRun = newClient.handlers.MakeBinary(&commands.UserBinaryData)
	commands.BinaryCmd.Run = newClient.handlers.AddData(&commands.DataTitle, &commands.Metadata, &commands.UserBinaryData)
	newClient.RootCmd.AddCommand(commands.AddDataCmd)

	commands.EditCredentialCmd.Run = newClient.handlers.EditData(&commands.EditedLocalID, &commands.EditedDataTitle, &commands.EditedMetadata, &commands.EditedUserCredential)
	commands.EditTextCmd.Run = newClient.handlers.EditData(&commands.EditedLocalID, &commands.EditedDataTitle, &commands.EditedMetadata, &commands.EditedUserText)
	commands.EditBankCardCmd.Run = newClient.handlers.EditData(&commands.EditedLocalID, &commands.EditedDataTitle, &commands.EditedMetadata, &commands.EditedUserBankCard)
	commands.EditBinaryCmd.PreRun = newClient.handlers.MakeBinary(&commands.EditedUserBinaryData)
	commands.EditBinaryCmd.Run = newClient.handlers.EditData(&commands.EditedLocalID, &commands.EditedDataTitle, &commands.EditedMetadata, &commands.EditedUserBinaryData)
	newClient.RootCmd.AddCommand(commands.EditDataCmd)

	commands.SendToServerCmd.Run = newClient.handlers.SendData(&commands.Masterkey)
	newClient.RootCmd.AddCommand(commands.SendToServerCmd)

	commands.GetFromServerCmd.Run = newClient.handlers.DownloadData(&commands.Masterkey)
	newClient.RootCmd.AddCommand(commands.GetFromServerCmd)

	return &newClient
}

func (c Client) CloseConnection() {
	c.GRPCConnect.Close()
}
