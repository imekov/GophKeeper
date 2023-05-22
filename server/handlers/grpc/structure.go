package grpc

import (
	pb "gophkeeper/proto"
	"gophkeeper/server/auth"
	"gophkeeper/server/storage/interfaces"
)

// GophKeeperServer поддерживает все необходимые методы сервера.
type GophKeeperServer struct {
	UserWriter interfaces.UserWriter
	DataWriter interfaces.DataWriter
	JWT        auth.JWT
	pb.UnimplementedGophKeeperServer
}
