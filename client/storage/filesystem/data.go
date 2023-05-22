package filesystem

import (
	"errors"

	"gophkeeper/client/storage/model"
)

// GetUserData возвращает пользовательские данные.
func (s FS) GetUserData() model.UserSession {
	return s.UserSession
}

// GetNewLocalID возвращает новые локальный идентификатор.
func (s FS) GetNewLocalID() int {
	return len(s.UserSession.DataArray) + 1
}

// GetUserDataByLocalID возвращает пользовательские данные по локальному идентификатору.
func (s FS) GetUserDataByLocalID(localID string) (resp model.UserData) {

	for i := range s.UserSession.DataArray {
		if s.UserSession.DataArray[i].LocalID == localID {
			return s.UserSession.DataArray[i]
		}
	}

	return
}

// AddData добавляет новые пользовательские данные в локальное хранилище.
func (s FS) AddData(newData []model.UserData) {
	for _, v := range newData {
		s.UserSession.DataArray = append(s.UserSession.DataArray, v)
	}
	s.SaveFile()
}

// UpdateDataByServerID обновляет данные по серверному идентификатору.
func (s FS) UpdateDataByServerID(updateData []model.UserData) {
	for _, v := range updateData {
		for k, m := range s.UserSession.DataArray {
			if m.ID == v.ID {
				s.UserSession.DataArray[k].Title = v.Title
				s.UserSession.DataArray[k].Metadata = v.Metadata
				s.UserSession.DataArray[k].Checksum = v.Checksum
				s.UserSession.DataArray[k].Data = v.Data
				break
			}
		}
	}

	s.SaveFile()
}

// UpdateDataByLocalID обновляет данные по локальному идентификатору.
func (s FS) UpdateDataByLocalID(updateData model.UserData) error {
	errorType := errors.New("тип сохраняемых данных не соответствует типу уже сохранённых")

	for k, m := range s.UserSession.DataArray {
		if m.LocalID == updateData.LocalID {

			switch newData := updateData.Data.(type) {

			case *model.Text:
				tempText := model.Text{}

				if oldTextData, ok := s.UserSession.DataArray[k].Data.(model.Text); ok {
					if newData.TextData != "" {
						tempText.TextData = newData.TextData
					} else {
						tempText.TextData = oldTextData.TextData
					}

				} else {
					return errorType
				}

				s.UserSession.DataArray[k].Data = tempText

			case *model.LoginPassword:
				tempLoginPassword := model.LoginPassword{}

				if oldLoginPasswordData, ok := s.UserSession.DataArray[k].Data.(model.LoginPassword); ok {
					if newData.Login != "" {
						tempLoginPassword.Login = newData.Login
					} else {
						tempLoginPassword.Login = oldLoginPasswordData.Login
					}

					if newData.Password != "" {
						tempLoginPassword.Password = newData.Password
					} else {
						tempLoginPassword.Password = oldLoginPasswordData.Password
					}
				} else {
					return errorType
				}

				s.UserSession.DataArray[k].Data = tempLoginPassword

			case *model.Binary:

				tempBinary := model.Binary{}

				if oldBinaryData, ok := s.UserSession.DataArray[k].Data.(model.Binary); ok {

					if newData.Path != "" {
						tempBinary.Path = newData.Path
						tempBinary.BinaryData = newData.BinaryData
					} else {
						tempBinary.Path = oldBinaryData.Path
						tempBinary.BinaryData = oldBinaryData.BinaryData
					}

				} else {
					return errorType
				}

				s.UserSession.DataArray[k].Data = tempBinary

			case *model.BankCard:
				tempBankCard := model.BankCard{}

				if oldBankCardData, ok := s.UserSession.DataArray[k].Data.(model.BankCard); ok {

					if newData.Number != "" {
						tempBankCard.Number = newData.Number
					} else {
						tempBankCard.Number = oldBankCardData.Number
					}

					if newData.ExpDate != "" {
						tempBankCard.ExpDate = newData.ExpDate
					} else {
						tempBankCard.ExpDate = oldBankCardData.ExpDate
					}

					if newData.Owner != "" {
						tempBankCard.Owner = newData.Owner
					} else {
						tempBankCard.Owner = oldBankCardData.Owner
					}

				} else {
					return errorType
				}

				s.UserSession.DataArray[k].Data = tempBankCard

			default:
				return errorType
			}

			if updateData.Title != "" {
				s.UserSession.DataArray[k].Title = updateData.Title
			}

			if updateData.Metadata != "" {
				s.UserSession.DataArray[k].Metadata = updateData.Metadata
			}

			s.UserSession.DataArray[k].Checksum = ""

			break
		}
	}

	s.SaveFile()
	return nil
}

// UpdateDataIDChecksumFromServer обновляет серверный идентификатор и контрольную сумму.
func (s FS) UpdateDataIDChecksumFromServer(localID string, ID string, checksum string) {
	for k := range s.UserSession.DataArray {
		if s.UserSession.DataArray[k].LocalID == localID {
			if ID != "" {
				s.UserSession.DataArray[k].ID = ID
			}

			s.UserSession.DataArray[k].Checksum = checksum
		}
	}
	s.SaveFile()
}

// IsDataExistByID проверяет существование данных по серверному идентификатору.
func (s FS) IsDataExistByID(dataID string) bool {
	for i := range s.UserSession.DataArray {
		if s.UserSession.DataArray[i].ID == dataID {
			return true
		}
	}
	return false
}
