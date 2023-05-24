package interfaces

import (
	"context"
	pb "gophkeeper/proto"
)

type DataWriter interface {
	InsertDataIntoDataTable(ctx context.Context, userID int, userData *pb.SendDataRequestArray) (dataID int, error error)
	UpdateData(ctx context.Context, userData *pb.UpdateDataRequestArray) (error error)
	GetUpdatesByChecksums(ctx context.Context, userID int, checkSums []string) ([]*pb.GetDataResponseArray, error)
}
