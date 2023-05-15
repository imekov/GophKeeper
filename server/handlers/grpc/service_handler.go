package grpc

import (
	"context"
	pb "gophkeeper/proto"
	"strconv"
)

func (s *GophKeeperServer) SendDataToServer(ctx context.Context, req *pb.SendDataRequest) (*pb.SendDataResponse, error) {
	var response pb.SendDataResponse
	userID, err := s.JWT.Validate(req.Token)
	if err != nil {
		return nil, err
	}

	for _, v := range req.Data {
		var respData pb.SendDataResponseArray

		dataID, err := s.Storage.InsertDataIntoDataTable(ctx, userID, v)
		if err != nil {
			return nil, err
		}
		respData.Checksum = v.Checksum
		respData.Id = strconv.Itoa(dataID)
		respData.LocalID = v.LocalID
		response.Resp = append(response.Resp, &respData)
	}

	return &response, nil
}

func (s *GophKeeperServer) GetDataFromServer(ctx context.Context, req *pb.GetDataRequest) (*pb.GetDataResponse, error) {
	var response pb.GetDataResponse
	userID, err := s.JWT.Validate(req.Token)
	if err != nil {
		return nil, err
	}

	response.Data, err = s.Storage.GetUpdatesByChecksums(ctx, userID, req.Checksum)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
