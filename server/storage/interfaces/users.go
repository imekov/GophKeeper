package interfaces

import "context"

type Users interface {
	CreateUser(ctx context.Context, login string, password string) (userID int, err error)
	LoginUser(ctx context.Context, login string, password string) (userID int, error error)
	IsUserExistByUserID(ctx context.Context, userID int) (response bool)
	IsUserExistByLogin(ctx context.Context, login string) (response bool)
}
