package grpc

import (
	"context"
	"errors"
	pb "gophkeeper/proto"
	"time"
)

func (s *GophKeeperServer) Register(ctx context.Context, req *pb.AuthRequest) (*pb.AuthResponse, error) {
	var response pb.AuthResponse

	if req.Login == "" || req.Password == "" {
		return nil, errors.New("пришли пустые данные")
	}

	if s.Storage.IsUserExistByLogin(ctx, req.Login) {
		return nil, errors.New("данный пользователь уже зарегистрирован")
	}

	userID, err := s.Storage.CreateUser(ctx, req.Login, req.Password)
	if err != nil {
		return nil, err
	}

	response.Token, err = s.JWT.Create(time.Hour*24, userID)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (s *GophKeeperServer) Login(ctx context.Context, req *pb.AuthRequest) (*pb.AuthResponse, error) {
	var response pb.AuthResponse

	if req.Login == "" || req.Password == "" {
		return nil, errors.New("пришли пустые данные")
	}

	if !s.Storage.IsUserExistByLogin(ctx, req.Login) {
		return nil, errors.New("неверные логин/password")
	}

	userID, err := s.Storage.LoginUser(ctx, req.Login, req.Password)
	if err != nil {
		return nil, err
	}

	response.Token, err = s.JWT.Create(time.Hour*24, userID)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
