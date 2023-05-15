package filesystem

import (
	"fmt"
	"gophkeeper/client/storage/model"
	"reflect"
	"strconv"
)

func (s FS) GetUserData() model.UserSession {
	return s.Data
}

func (s FS) GetNewLocalID() string {
	return strconv.Itoa(len(s.Data.Data) + 1)
}

func (s FS) GetUserDataByLocalID(localID string) (resp model.UserData) {

	for i := range s.Data.Data {
		if s.Data.Data[i].LocalID == localID {
			return s.Data.Data[i]
		}
	}

	return
}

func (s FS) AddData(newData []model.UserData) {
	for _, v := range newData {
		s.Data.Data = append(s.Data.Data, v)
	}
	s.SaveFile()
}

func (s FS) UpdateDataByServerID(updateData []model.UserData) {
	for _, v := range updateData {
		for k, m := range s.Data.Data {
			if m.ID == v.ID {
				s.Data.Data[k].Title = v.Title
				s.Data.Data[k].Metadata = v.Metadata
				s.Data.Data[k].Checksum = v.Checksum
				s.Data.Data[k].Data = v.Data
				break
			}
		}
	}

	s.SaveFile()
}

func (s FS) UpdateDataByLocalID(updateData model.UserData) {
	for k, m := range s.Data.Data {
		if m.LocalID == updateData.LocalID {
			if reflect.TypeOf(s.Data.Data[k].Data) == reflect.TypeOf(updateData.Data) {
				if updateData.Title != "" {
					s.Data.Data[k].Title = updateData.Title
				}

				if updateData.Metadata != "" {
					s.Data.Data[k].Metadata = updateData.Metadata
				}

				if updateData.Checksum != "" {
					s.Data.Data[k].Checksum = updateData.Checksum
				}

				s.Data.Data[k].Data = updateData.Data

			} else {
				fmt.Println("Не удалось изменить информацию, так как тип сохранненых данных отличается от типа введённых.")
			}

			break
		}
	}

	s.SaveFile()
}

func (s FS) UpdateCredentialByLocalID(updateData model.UserData) {
	for k, m := range s.Data.Data {
		if m.LocalID == updateData.LocalID {
			if reflect.TypeOf(s.Data.Data[k].Data) == reflect.TypeOf(updateData.Data) {
				if updateData.Title != "" {
					s.Data.Data[k].Title = updateData.Title
				}

				if updateData.Metadata != "" {
					s.Data.Data[k].Metadata = updateData.Metadata
				}

				if updateData.Checksum != "" {
					s.Data.Data[k].Checksum = updateData.Checksum
				}

				s.Data.Data[k].Data = updateData.Data

			} else {
				fmt.Println("Не удалось изменить информацию, так как тип сохранненых данных отличается от типа введённых.")
			}

			break
		}
	}

	s.SaveFile()
}

func (s FS) UpdateDataIDChecksumFromServer(localID string, ID string, checksum string) {
	for k := range s.Data.Data {
		if s.Data.Data[k].LocalID == localID {
			s.Data.Data[k].ID = ID
			s.Data.Data[k].Checksum = checksum
		}
	}
	s.SaveFile()
}

func (s FS) IsDataExistByID(dataID string) bool {
	for i := range s.Data.Data {
		if s.Data.Data[i].ID == dataID {
			return true
		}
	}
	return false
}
