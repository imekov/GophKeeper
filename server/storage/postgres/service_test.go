package postgres

import (
	"context"
	"database/sql"
	"reflect"
	"testing"

	"github.com/rs/zerolog"

	pb "gophkeeper/proto"
)

func TestDBConnect_CreateUser(t *testing.T) {
	type fields struct {
		DBConnect *sql.DB
		Logger    *zerolog.Logger
	}
	type args struct {
		ctx      context.Context
		login    string
		password string
	}
	var tests []struct {
		name       string
		fields     fields
		args       args
		wantUserID int
		wantErr    bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := DBConnect{
				DBConnect: tt.fields.DBConnect,
				Logger:    tt.fields.Logger,
			}
			gotUserID, err := s.CreateUser(tt.args.ctx, tt.args.login, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotUserID != tt.wantUserID {
				t.Errorf("CreateUser() gotUserID = %v, want %v", gotUserID, tt.wantUserID)
			}
		})
	}
}

func TestDBConnect_GetUpdatesByChecksums(t *testing.T) {
	type fields struct {
		DBConnect *sql.DB
		Logger    *zerolog.Logger
	}
	type args struct {
		ctx       context.Context
		userID    int
		checkSums []string
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		want    []*pb.GetDataResponseArray
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := DBConnect{
				DBConnect: tt.fields.DBConnect,
				Logger:    tt.fields.Logger,
			}
			got, err := s.GetUpdatesByChecksums(tt.args.ctx, tt.args.userID, tt.args.checkSums)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUpdatesByChecksums() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUpdatesByChecksums() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDBConnect_InsertDataIntoDataTable(t *testing.T) {
	type fields struct {
		DBConnect *sql.DB
		Logger    *zerolog.Logger
	}
	type args struct {
		ctx      context.Context
		userID   int
		userData *pb.SendDataRequestArray
	}
	var tests []struct {
		name       string
		fields     fields
		args       args
		wantDataID int
		wantErr    bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := DBConnect{
				DBConnect: tt.fields.DBConnect,
				Logger:    tt.fields.Logger,
			}
			gotDataID, err := s.InsertDataIntoDataTable(tt.args.ctx, tt.args.userID, tt.args.userData)
			if (err != nil) != tt.wantErr {
				t.Errorf("InsertDataIntoDataTable() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotDataID != tt.wantDataID {
				t.Errorf("InsertDataIntoDataTable() gotDataID = %v, want %v", gotDataID, tt.wantDataID)
			}
		})
	}
}

func TestDBConnect_IsUserExistByLogin(t *testing.T) {
	type fields struct {
		DBConnect *sql.DB
		Logger    *zerolog.Logger
	}
	type args struct {
		ctx   context.Context
		login string
	}
	var tests []struct {
		name         string
		fields       fields
		args         args
		wantResponse bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := DBConnect{
				DBConnect: tt.fields.DBConnect,
				Logger:    tt.fields.Logger,
			}
			if gotResponse := s.IsUserExistByLogin(tt.args.ctx, tt.args.login); gotResponse != tt.wantResponse {
				t.Errorf("IsUserExistByLogin() = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func TestDBConnect_IsUserExistByUserID(t *testing.T) {
	type fields struct {
		DBConnect *sql.DB
		Logger    *zerolog.Logger
	}
	type args struct {
		ctx    context.Context
		userID int
	}
	var tests []struct {
		name         string
		fields       fields
		args         args
		wantResponse bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := DBConnect{
				DBConnect: tt.fields.DBConnect,
				Logger:    tt.fields.Logger,
			}
			if gotResponse := s.IsUserExistByUserID(tt.args.ctx, tt.args.userID); gotResponse != tt.wantResponse {
				t.Errorf("IsUserExistByUserID() = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func TestDBConnect_LoginUser(t *testing.T) {
	type fields struct {
		DBConnect *sql.DB
		Logger    *zerolog.Logger
	}
	type args struct {
		ctx      context.Context
		login    string
		password string
	}
	var tests []struct {
		name       string
		fields     fields
		args       args
		wantUserID int
		wantErr    bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := DBConnect{
				DBConnect: tt.fields.DBConnect,
				Logger:    tt.fields.Logger,
			}
			gotUserID, err := s.LoginUser(tt.args.ctx, tt.args.login, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoginUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotUserID != tt.wantUserID {
				t.Errorf("LoginUser() gotUserID = %v, want %v", gotUserID, tt.wantUserID)
			}
		})
	}
}

func TestDBConnect_UpdateData(t *testing.T) {
	type fields struct {
		DBConnect *sql.DB
		Logger    *zerolog.Logger
	}
	type args struct {
		ctx      context.Context
		userData *pb.UpdateDataRequestArray
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := DBConnect{
				DBConnect: tt.fields.DBConnect,
				Logger:    tt.fields.Logger,
			}
			if err := s.UpdateData(tt.args.ctx, tt.args.userData); (err != nil) != tt.wantErr {
				t.Errorf("UpdateData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetNewConnection(t *testing.T) {
	type args struct {
		db     *sql.DB
		dbConf string
		logger *zerolog.Logger
	}
	var tests []struct {
		name string
		args args
		want *DBConnect
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetNewConnection(tt.args.db, tt.args.dbConf, tt.args.logger); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetNewConnection() = %v, want %v", got, tt.want)
			}
		})
	}
}
