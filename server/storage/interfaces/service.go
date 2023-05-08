package interfaces

import "context"

type Service interface {
	InsertDataIntoDataTable(data []byte, userID int, dataType string, checksum string, ctx context.Context) (error error)
	GetUpdates(userID int, userDataID []string, ctx context.Context) (data []byte)
}
