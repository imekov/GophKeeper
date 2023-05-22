package grpc

import (
	"context"
	"reflect"
	"testing"

	pb "gophkeeper/proto"
	"gophkeeper/server/auth"
	"gophkeeper/server/storage/interfaces"
)

func TestGophKeeperServer_GetDataFromServer(t *testing.T) {
	type fields struct {
		UserWriter                    interfaces.UserWriter
		DataWriter                    interfaces.DataWriter
		JWT                           auth.JWT
		UnimplementedGophKeeperServer pb.UnimplementedGophKeeperServer
	}
	type args struct {
		ctx context.Context
		req *pb.GetDataRequest
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		want    *pb.GetDataResponse
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &GophKeeperServer{
				UserWriter:                    tt.fields.UserWriter,
				DataWriter:                    tt.fields.DataWriter,
				JWT:                           tt.fields.JWT,
				UnimplementedGophKeeperServer: tt.fields.UnimplementedGophKeeperServer,
			}
			got, err := s.GetDataFromServer(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDataFromServer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDataFromServer() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGophKeeperServer_Login(t *testing.T) {
	type fields struct {
		UserWriter                    interfaces.UserWriter
		DataWriter                    interfaces.DataWriter
		JWT                           auth.JWT
		UnimplementedGophKeeperServer pb.UnimplementedGophKeeperServer
	}
	type args struct {
		ctx context.Context
		req *pb.AuthRequest
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		want    *pb.AuthResponse
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &GophKeeperServer{
				UserWriter:                    tt.fields.UserWriter,
				DataWriter:                    tt.fields.DataWriter,
				JWT:                           tt.fields.JWT,
				UnimplementedGophKeeperServer: tt.fields.UnimplementedGophKeeperServer,
			}
			got, err := s.Login(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Login() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGophKeeperServer_Register(t *testing.T) {
	type fields struct {
		UserWriter                    interfaces.UserWriter
		DataWriter                    interfaces.DataWriter
		JWT                           auth.JWT
		UnimplementedGophKeeperServer pb.UnimplementedGophKeeperServer
	}
	type args struct {
		ctx context.Context
		req *pb.AuthRequest
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		want    *pb.AuthResponse
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &GophKeeperServer{
				UserWriter:                    tt.fields.UserWriter,
				DataWriter:                    tt.fields.DataWriter,
				JWT:                           tt.fields.JWT,
				UnimplementedGophKeeperServer: tt.fields.UnimplementedGophKeeperServer,
			}
			got, err := s.Register(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Register() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGophKeeperServer_SendDataToServer(t *testing.T) {
	type fields struct {
		UserWriter                    interfaces.UserWriter
		DataWriter                    interfaces.DataWriter
		JWT                           auth.JWT
		UnimplementedGophKeeperServer pb.UnimplementedGophKeeperServer
	}
	type args struct {
		ctx context.Context
		req *pb.SendDataRequest
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		want    *pb.SendDataResponse
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &GophKeeperServer{
				UserWriter:                    tt.fields.UserWriter,
				DataWriter:                    tt.fields.DataWriter,
				JWT:                           tt.fields.JWT,
				UnimplementedGophKeeperServer: tt.fields.UnimplementedGophKeeperServer,
			}
			got, err := s.SendDataToServer(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("SendDataToServer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SendDataToServer() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGophKeeperServer_UpdateDataToServer(t *testing.T) {
	type fields struct {
		UserWriter                    interfaces.UserWriter
		DataWriter                    interfaces.DataWriter
		JWT                           auth.JWT
		UnimplementedGophKeeperServer pb.UnimplementedGophKeeperServer
	}
	type args struct {
		ctx context.Context
		req *pb.UpdateDataRequest
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		want    *pb.UpdateDataResponse
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &GophKeeperServer{
				UserWriter:                    tt.fields.UserWriter,
				DataWriter:                    tt.fields.DataWriter,
				JWT:                           tt.fields.JWT,
				UnimplementedGophKeeperServer: tt.fields.UnimplementedGophKeeperServer,
			}
			got, err := s.UpdateDataToServer(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateDataToServer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateDataToServer() got = %v, want %v", got, tt.want)
			}
		})
	}
}
