package server

import (
	"database/sql"
	"fmt"
	"github.com/rs/zerolog"
	pb "gophkeeper/proto"
	"gophkeeper/server/auth"
	grpc_handler "gophkeeper/server/handlers/grpc"
	"gophkeeper/server/storage/postgres"
	"net"
	"time"

	"google.golang.org/grpc"
)

type Server struct {
	dbConnection    *sql.DB
	postgresConnect *postgres.DBConnect
}

func DBConnectionInitialization(DatabaseURI string, logger *zerolog.Logger) (s Server) {
	var err error

	for i := 0; i <= 10; i++ {
		s.dbConnection, err = sql.Open("postgres", DatabaseURI)
		if err == nil {
			break
		}
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		newPrint := fmt.Sprintf("unable to connect to database %v\n: %s", DatabaseURI, err.Error())
		logger.Error().Msg(newPrint)
	}

	s.postgresConnect = postgres.GetNewConnection(s.dbConnection, DatabaseURI, logger)

	return
}

func (s Server) DBCloseConnection() {
	s.dbConnection.Close()
}

func (s Server) StartGRPCServer(runAddress string, logger *zerolog.Logger) error {
	listen, err := net.Listen("tcp", runAddress)
	if err != nil {
		newPrint := fmt.Sprintf("unable to create new tcp listener : %s", err.Error())
		logger.Error().Msg(newPrint)
	}

	newSrv := grpc.NewServer()
	pb.RegisterGophKeeperServer(newSrv, &grpc_handler.GophKeeperServer{UserWriter: s.postgresConnect, DataWriter: s.postgresConnect, JWT: auth.NewJWT(logger)})
	fmt.Printf("Сервер gRPC начал работу на порту: %s", runAddress)

	return newSrv.Serve(listen)
}
