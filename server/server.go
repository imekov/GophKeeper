package server

import (
	"database/sql"
	"fmt"
	pb "gophkeeper/proto"
	"gophkeeper/server/auth"
	grpc_handler "gophkeeper/server/handlers/grpc"
	"gophkeeper/server/storage/postgres"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

type Server struct {
	dbConnection    *sql.DB
	postgresConnect *postgres.DBConnect
}

func DBConnectionInitialization(DatabaseURI string) (s Server) {
	var err error

	for i := 0; i <= 10; i++ {
		s.dbConnection, err = sql.Open("postgres", DatabaseURI)
		if err == nil {
			break
		}
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		log.Fatalf("unable to connect to database %v\n", DatabaseURI)
	}

	s.postgresConnect = postgres.GetNewConnection(s.dbConnection, DatabaseURI)

	return
}

func (s Server) DBCloseConnection() {
	s.dbConnection.Close()
}

func (s Server) StartGRPCServer(runAddress string) error {
	listen, err := net.Listen("tcp", runAddress)
	if err != nil {
		log.Fatal(err)
	}

	newSrv := grpc.NewServer()
	pb.RegisterGophKeeperServer(newSrv, &grpc_handler.GophKeeperServer{Storage: s.postgresConnect, JWT: auth.NewJWT()})
	fmt.Printf("Сервер gRPC начал работу на порту: %s", runAddress)

	return newSrv.Serve(listen)
}
