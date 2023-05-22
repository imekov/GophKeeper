package handlers

import (
	"reflect"
	"testing"

	"github.com/spf13/cobra"

	"gophkeeper/client/storage"
	"gophkeeper/client/storage/model"
	"gophkeeper/proto"
)

func TestHandlers_AddData(t *testing.T) {
	type fields struct {
		Client proto.GophKeeperClient
		Repo   storage.Repo
	}
	type args struct {
		title    *string
		metadata *string
		data     any
	}
	var tests []struct {
		name   string
		fields fields
		args   args
		want   func(cmd *cobra.Command, args []string)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := Handlers{
				Client: tt.fields.Client,
				Repo:   tt.fields.Repo,
			}
			if got := h.AddData(tt.args.title, tt.args.metadata, tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandlers_EditData(t *testing.T) {
	type fields struct {
		Client proto.GophKeeperClient
		Repo   storage.Repo
	}
	type args struct {
		localID  *string
		title    *string
		metadata *string
		data     any
	}
	var tests []struct {
		name   string
		fields fields
		args   args
		want   func(cmd *cobra.Command, args []string)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := Handlers{
				Client: tt.fields.Client,
				Repo:   tt.fields.Repo,
			}
			if got := h.EditData(tt.args.localID, tt.args.title, tt.args.metadata, tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EditData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandlers_GetUserData(t *testing.T) {
	type fields struct {
		Client proto.GophKeeperClient
		Repo   storage.Repo
	}
	type args struct {
		index *string
	}
	var tests []struct {
		name   string
		fields fields
		args   args
		want   func(cmd *cobra.Command, args []string)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := Handlers{
				Client: tt.fields.Client,
				Repo:   tt.fields.Repo,
			}
			if got := h.GetUserData(tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandlers_GetUserDataList(t *testing.T) {
	type fields struct {
		Client proto.GophKeeperClient
		Repo   storage.Repo
	}
	var tests []struct {
		name   string
		fields fields
		want   func(cmd *cobra.Command, args []string)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := Handlers{
				Client: tt.fields.Client,
				Repo:   tt.fields.Repo,
			}
			if got := h.GetUserDataList(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserDataList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandlers_Login(t *testing.T) {
	type fields struct {
		Client proto.GophKeeperClient
		Repo   storage.Repo
	}
	type args struct {
		login    *string
		password *string
	}
	var tests []struct {
		name   string
		fields fields
		args   args
		want   func(cmd *cobra.Command, args []string)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := Handlers{
				Client: tt.fields.Client,
				Repo:   tt.fields.Repo,
			}
			if got := h.Login(tt.args.login, tt.args.password); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Login() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandlers_MakeBinary(t *testing.T) {
	type fields struct {
		Client proto.GophKeeperClient
		Repo   storage.Repo
	}
	type args struct {
		binary *model.Binary
	}
	var tests []struct {
		name   string
		fields fields
		args   args
		want   func(cmd *cobra.Command, args []string)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := Handlers{
				Client: tt.fields.Client,
				Repo:   tt.fields.Repo,
			}
			if got := h.MakeBinary(tt.args.binary); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandlers_Register(t *testing.T) {
	type fields struct {
		Client proto.GophKeeperClient
		Repo   storage.Repo
	}
	type args struct {
		login    *string
		password *string
	}
	var tests []struct {
		name   string
		fields fields
		args   args
		want   func(cmd *cobra.Command, args []string)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := Handlers{
				Client: tt.fields.Client,
				Repo:   tt.fields.Repo,
			}
			if got := h.Register(tt.args.login, tt.args.password); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Register() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandlers_Sync(t *testing.T) {
	type fields struct {
		Client proto.GophKeeperClient
		Repo   storage.Repo
	}
	type args struct {
		masterKey *string
	}
	var tests []struct {
		name   string
		fields fields
		args   args
		want   func(cmd *cobra.Command, args []string)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := Handlers{
				Client: tt.fields.Client,
				Repo:   tt.fields.Repo,
			}
			if got := h.Sync(tt.args.masterKey); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sync() = %v, want %v", got, tt.want)
			}
		})
	}
}
