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

type Handlers struct {
	Client pb.GophKeeperClient
	Repo   storage.Repo
}
