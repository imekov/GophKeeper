package storage

import (
	"gophkeeper/client/storage/model"
	"os"
)

type Repo interface {
	IsDataExistByID(dataID string) bool
	GetUserData() model.UserSession
	GetUserDataByLocalID(localID string) (resp model.UserData)
	AddData(newData []model.UserData)
	UpdateDataByServerID(updateData []model.UserData)
	UpdateDataByLocalID(updateData model.UserData)
	UpdateDataIDChecksumFromServer(localID string, ID string, checksum string)
	GetNewLocalID() string

	IsUserAuthorized() bool
	UpdateToken(newToken string)

	OpenFile(flag int) *os.File
	ReadFile(data *model.UserSession) error
	SaveFile() error
}
