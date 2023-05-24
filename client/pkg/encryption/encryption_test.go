package encryption

import (
	"reflect"
	"testing"
)

func TestDecode(t *testing.T) {
	type args struct {
		encryptData []byte
		masterKey   string
	}
	var tests []struct {
		name     string
		args     args
		wantResp any
		wantErr  bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := Decode(tt.args.encryptData, tt.args.masterKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("Decode() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestDecrypt(t *testing.T) {
	type args struct {
		encryptData []byte
		masterKey   string
	}
	var tests []struct {
		name     string
		args     args
		wantResp []byte
		wantErr  bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := Decrypt(tt.args.encryptData, tt.args.masterKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("Decrypt() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestEncode(t *testing.T) {
	type args struct {
		data      any
		masterkey string
	}
	var tests []struct {
		name     string
		args     args
		wantResp []byte
		wantErr  bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := Encode(tt.args.data, tt.args.masterkey)
			if (err != nil) != tt.wantErr {
				t.Errorf("Encode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("Encode() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func Test_getHexKey(t *testing.T) {
	type args struct {
		key string
	}
	var tests []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getHexKey(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("getHexKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getHexKey() got = %v, want %v", got, tt.want)
			}
		})
	}
}
