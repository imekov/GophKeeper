package handlers

import (
	"context"
	"crypto/md5"
	"fmt"
	"github.com/spf13/cobra"
	"gophkeeper/client/pkg/encryption"
	"gophkeeper/client/storage/model"
	pb "gophkeeper/proto"
)

// SendData отправляет новые данные на сервер. Отсутствие серверного ID считается признаком
// данных, которые существуют только в локальном хранилище.
func (h Handlers) SendData(masterKey *string) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		if h.Repo.IsUserAuthorized() {
			data := h.Repo.GetUserData()
			req := pb.SendDataRequest{Token: data.Token}

			for _, v := range data.Data {
				if v.ID == "" {
					var dataType string
					enc, err := encryption.Encode(v.Data, *masterKey)
					if err != nil {
						fmt.Printf("Возникла ошибка шифрования при отправке данных:\n%s", err.Error())
						return
					}

					switch v.Data.(type) {
					case model.BankCard:
						dataType = BANKCARDTYPE
					case model.Binary:
						dataType = BINARYTYPE
					case model.Text:
						dataType = TEXTTYPE
					case model.LoginPassword:
						dataType = LOGINTYPE
					}

					c := fmt.Sprintf("%x", md5.Sum(enc))
					s := pb.SendDataRequestArray{
						LocalID:  v.LocalID,
						Title:    v.Title,
						Metadata: v.Metadata,
						Checksum: c,
						DataType: dataType,
						Data:     enc,
					}
					req.Data = append(req.Data, &s)
				}
			}

			if len(req.Data) == 0 {
				fmt.Println("Нет новых данных для отправки на сервер.")
			} else {
				resp, err := h.Client.SendDataToServer(context.Background(), &req)
				if err != nil {
					fmt.Printf("Возникла ошибка при отправке данных на сервер:\n%s", err.Error())
					return
				}

				fmt.Println("Данные отправлены на сервер. Присвоены следующие ID:")
				for k, v := range resp.Resp {
					fmt.Printf("%d. %s\n", k+1, v.Id)
					h.Repo.UpdateDataIDChecksumFromServer(v.LocalID, v.Id, v.Checksum)
				}
			}

		} else {
			fmt.Println("Для отправки данных нужно авторизоваться.")
		}

	}
}

// DownloadData скачивает данные с сервера. Отправляет массив с контрольными суммами для
// получения новых данных и изменений по существующим записям.
func (h Handlers) DownloadData(masterKey *string) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		if h.Repo.IsUserAuthorized() {
			data := h.Repo.GetUserData()
			req := pb.GetDataRequest{Token: data.Token, Checksum: []string{""}}

			for _, v := range data.Data {
				if v.ID != "" && v.Checksum != "" {
					req.Checksum = append(req.Checksum, v.Checksum)
				}
			}

			resp, err := h.Client.GetDataFromServer(context.Background(), &req)
			if err != nil {
				fmt.Printf("Возникла ошибка при скачивании данных с сервера:\n%s", err.Error())
				return
			}

			var dataToSave []model.UserData
			var dataToUpdate []model.UserData

			for _, v := range resp.Data {
				s := model.UserData{
					ID:       v.DataID,
					Checksum: v.CheckSum,
					Title:    v.Title,
					Metadata: v.Metadata,
				}

				switch v.DataType {
				case LOGINTYPE:
					s.Data, err = encryption.Decode[model.LoginPassword](v.Userdata, *masterKey)
				case BINARYTYPE:
					s.Data, err = encryption.Decode[model.Binary](v.Userdata, *masterKey)
				case BANKCARDTYPE:
					s.Data, err = encryption.Decode[model.BankCard](v.Userdata, *masterKey)
				case TEXTTYPE:
					s.Data, err = encryption.Decode[model.Text](v.Userdata, *masterKey)
				}

				if err != nil {
					fmt.Printf("Возникла ошибка дешифровки данных после получения с сервера:\n%s", err.Error())
					return
				}

				if h.Repo.IsDataExistByID(v.DataID) {
					dataToUpdate = append(dataToUpdate, s)
				} else {
					s.LocalID = h.Repo.GetNewLocalID()
					dataToSave = append(dataToSave, s)
				}

			}

			h.Repo.UpdateDataByServerID(dataToUpdate)
			h.Repo.AddData(dataToSave)

			fmt.Println("Данные загружены.")

		} else {
			fmt.Println("Для получения данных нужно авторизоваться.")
		}
	}
}
