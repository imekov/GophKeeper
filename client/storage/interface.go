package storage

import (
	"os"

	"gophkeeper/client/storage/model"
)

type Repo interface {
	IsDataExistByID(dataID string) bool
	GetUserData() model.UserSession
	GetUserDataByLocalID(localID string) (resp model.UserData)
	AddData(newData []model.UserData)
	UpdateDataByServerID(updateData []model.UserData)
	UpdateDataByLocalID(updateData model.UserData) error
	UpdateDataIDChecksumFromServer(localID string, ID string, checksum string)
	GetNewLocalID() int

	IsUserAuthorized() bool
	UpdateToken(newToken string)

	OpenFile(flag int) *os.File
	ReadFile(data *model.UserSession) error
	SaveFile() error
}
