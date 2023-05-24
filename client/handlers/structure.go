package handlers

import (
	"gophkeeper/client/storage"
	pb "gophkeeper/proto"
)

const (
	BANKCARDTYPE string = "BANKCARD"
	BINARYTYPE   string = "BINARY"
	TEXTTYPE     string = "TEXT"
	LOGINTYPE    string = "LOGIN"
)

// Handlers содержит ссылки на grpc клиент и интерфейс репозитория.
type Handlers struct {
	Client pb.GophKeeperClient
	Repo   storage.Repo
}

// request собирает все входящие данных для отправки запроса на обновление.
type request struct {
	Send   pb.SendDataRequest
	Update pb.UpdateDataRequest
	Get    pb.GetDataRequest
}
