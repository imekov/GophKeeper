package postgres

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// CreateUser создаёт нового пользователя в БД.
func (s DBConnect) CreateUser(ctx context.Context, login string, password string) (userID int, error error) {

	tx, err := s.DBConnect.BeginTx(ctx, nil)
	if err != nil {
		s.Logger.Error().Msg(err.Error())
		return 0, err
	}
	defer tx.Rollback()

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		s.Logger.Error().Msg(err.Error())
		return 0, err
	}

	err = tx.QueryRow("INSERT INTO users (login, password) VALUES ($1, $2) ON CONFLICT (login) DO NOTHING RETURNING user_ID;", login, hash).Scan(&userID)
	if err != nil {
		s.Logger.Error().Msg(err.Error())
		return 0, err
	}

	tx.Commit()

	return userID, nil
}

// LoginUser авторизация пользователя.
func (s DBConnect) LoginUser(ctx context.Context, login string, password string) (userID int, error error) {

	tx, err := s.DBConnect.BeginTx(ctx, nil)
	if err != nil {
		s.Logger.Error().Msg(err.Error())
		return 0, err
	}

	var hashPassword string

	err = tx.QueryRow("select user_ID, password from users where login = $1;", login).Scan(&userID, &hashPassword)
	if err != nil {
		s.Logger.Error().Msg(err.Error())
		return 0, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	if err != nil {
		s.Logger.Error().Msg(err.Error())
		return 0, errors.New("неверные логин/password")
	}

	return userID, nil
}

// IsUserExistByLogin проверяет, зарегистрирован ли пользователь по логину.
func (s DBConnect) IsUserExistByLogin(ctx context.Context, login string) (response bool) {
	response = false

	tx, err := s.DBConnect.BeginTx(ctx, nil)
	if err != nil {
		s.Logger.Error().Msg(err.Error())
	}

	var countOfRows int

	err = tx.QueryRow("select COUNT(*) from users where login = $1;", login).Scan(&countOfRows)
	if err != nil {
		s.Logger.Error().Msg(err.Error())
		return false
	}

	if countOfRows != 0 {
		response = true
	}

	return response
}

// IsUserExistByUserID проверяет, зарегистрирован ли пользователь по идентификатору.
func (s DBConnect) IsUserExistByUserID(ctx context.Context, userID int) (response bool) {
	response = false

	tx, err := s.DBConnect.BeginTx(ctx, nil)
	if err != nil {
		s.Logger.Error().Msg(err.Error())
	}

	var countOfRows int

	err = tx.QueryRow("select COUNT(*) from users where user_ID = $1;", userID).Scan(&countOfRows)
	if err != nil {
		s.Logger.Error().Msg(err.Error())
		return false
	}

	if countOfRows != 0 {
		response = true
	}

	return response
}
