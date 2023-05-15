package interfaces

import (
	"context"
	pb "gophkeeper/proto"
)

type Service interface {
	InsertDataIntoDataTable(ctx context.Context, userID int, userData *pb.SendDataRequestArray) (dataID int, error error)
	GetUpdatesByChecksums(ctx context.Context, userID int, checkSums []string) ([]*pb.GetDataResponseArray, error)
}
