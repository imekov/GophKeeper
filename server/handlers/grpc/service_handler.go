package grpc

import (
	"context"
	"strconv"

	pb "gophkeeper/proto"
)

// SendDataToServer принимает данные на сервер.
func (s *GophKeeperServer) SendDataToServer(ctx context.Context, req *pb.SendDataRequest) (*pb.SendDataResponse, error) {
	var response pb.SendDataResponse
	userID, err := s.JWT.Validate(req.Token)
	if err != nil {
		return nil, err
	}

	for _, v := range req.Data {
		var respData pb.SendDataResponseArray

		dataID, err := s.DataWriter.InsertDataIntoDataTable(ctx, userID, v)
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

// UpdateDataToServer принимает обновления на сервер.
func (s *GophKeeperServer) UpdateDataToServer(ctx context.Context, req *pb.UpdateDataRequest) (*pb.UpdateDataResponse, error) {
	var response pb.UpdateDataResponse
	_, err := s.JWT.Validate(req.Token)
	if err != nil {
		return nil, err
	}

	for _, v := range req.Data {
		var respData pb.UpdateDataResponseArray

		if err = s.DataWriter.UpdateData(ctx, v); err != nil {
			return nil, err
		}

		respData.Ok = true
		respData.CheckSum = v.Checksum
		respData.LocalID = v.LocalID
		response.Resp = append(response.Resp, &respData)
	}

	return &response, nil
}

// GetDataFromServer отправляет данные на клиент.
func (s *GophKeeperServer) GetDataFromServer(ctx context.Context, req *pb.GetDataRequest) (*pb.GetDataResponse, error) {
	var response pb.GetDataResponse
	userID, err := s.JWT.Validate(req.Token)
	if err != nil {
		return nil, err
	}

	response.Data, err = s.DataWriter.GetUpdatesByChecksums(ctx, userID, req.Checksum)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
