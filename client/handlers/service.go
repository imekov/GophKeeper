package handlers

import (
	"bufio"
	"context"
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	"gophkeeper/client/pkg/encryption"
	"gophkeeper/client/storage/model"
	pb "gophkeeper/proto"

	"github.com/spf13/cobra"
)

// AddData добавляет новые данные в локальное хранилище.
func (h Handlers) AddData(title *string, metadata *string, data any) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		localID := h.Repo.GetNewLocalID()
		h.Repo.AddData([]model.UserData{{
			LocalID:  strconv.Itoa(localID),
			Title:    *title,
			Metadata: *metadata,
			Data:     data,
		}})
		fmt.Println("данные успешно добавлены")
	}
}

// EditData производит изменение созданных данных. Первым параметром передаётся айди записи, далее новая информация.
func (h Handlers) EditData(localID *string, title *string, metadata *string, data any) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		if err := h.Repo.UpdateDataByLocalID(model.UserData{
			LocalID:  *localID,
			Title:    *title,
			Metadata: *metadata,
			Data:     data,
		}); err != nil {
			fmt.Printf("Невозможно обновить данные: %s", err.Error())
			return
		}

		fmt.Println("данные успешно изменены")
	}
}

// MakeBinary конвертирует файл в двоичные данные.
func (h Handlers) MakeBinary(binary *model.Binary) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		if binary.Path != "" {
			file, err := os.Open(binary.Path)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer file.Close()

			stat, err := file.Stat()
			if err != nil {
				fmt.Println(err)
				return
			}

			binary.BinaryData = make([]byte, stat.Size())
			_, err = bufio.NewReader(file).Read(binary.BinaryData)
			if err != nil && err != io.EOF {
				fmt.Println(err)
				return
			}
		}
	}
}

// GetUserDataList выводит список всех локальных данных.
func (h Handlers) GetUserDataList() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		data := h.Repo.GetUserData()
		if len(data.DataArray) == 0 {
			fmt.Println("нет данных для отображения")
		} else {
			fmt.Println("список сохранённых данных")
			for _, v := range data.DataArray {
				var dataType string
				switch v.Data.(type) {
				case model.Binary:
					dataType = "бинарные данные"
				case model.Text:
					dataType = "текстовые данные"
				case model.BankCard:
					dataType = "данные банковских карт"
				case model.LoginPassword:
					dataType = "логин и пароль"
				}

				fmt.Printf("%s. %s\nTitle:%s\nMetadata:%s\nID:%s\nChecksum: %s\n\n", v.LocalID, dataType, v.Title, v.Metadata, v.ID, v.Checksum)
			}
		}
	}
}

// GetUserData выводит информацию по выбранным данным.
func (h Handlers) GetUserData(index *string) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		data := h.Repo.GetUserDataByLocalID(*index)
		fmt.Printf("Title: %s\nMetadata: %s\n\nData\n%s", data.Title, data.Metadata, data.Data)
	}
}

// Sync производит синхронизацию данных.
func (h Handlers) Sync(masterKey *string) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		if h.Repo.IsUserAuthorized() {

			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			data := h.Repo.GetUserData()
			req := request{
				Send:   pb.SendDataRequest{Token: data.Token},
				Update: pb.UpdateDataRequest{Token: data.Token},
				Get:    pb.GetDataRequest{Token: data.Token, Checksum: []string{""}},
			}

			for _, v := range data.DataArray {

				if v.ID == "" { //send
					enc, err := encryption.Encode(v.Data, *masterKey)
					if err != nil {
						fmt.Printf("Возникла ошибка шифрования при отправке данных:\n%s", err.Error())
						return
					}

					var dataType string

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

					req.Send.Data = append(req.Send.Data, &s)

				} else {

					if v.Checksum == "" { //update
						enc, err := encryption.Encode(v.Data, *masterKey)
						if err != nil {
							fmt.Printf("Возникла ошибка шифрования при отправке обновлённых данных:\n%s", err.Error())
							return
						}

						c := fmt.Sprintf("%x", md5.Sum(enc))
						s := pb.UpdateDataRequestArray{
							ServerID: v.ID,
							LocalID:  v.LocalID,
							Title:    v.Title,
							Metadata: v.Metadata,
							Checksum: c,
							Data:     enc,
						}
						req.Update.Data = append(req.Update.Data, &s)

					} else { //get
						req.Get.Checksum = append(req.Get.Checksum, v.Checksum)
					}

				}

			}

			//Сначала отправляем обновления старых данных
			if len(req.Update.Data) == 0 {
				fmt.Println("Нет обновлённых данных для отправки на сервер.")
			} else {
				respUpdate, err := h.Client.UpdateDataToServer(ctx, &req.Update)
				if err != nil {
					fmt.Printf("Возникла ошибка при отправке обновлённых данных на сервер:\n%s", err.Error())
					return
				}

				fmt.Println("Обновленные данные отправлены на сервер.")
				for _, d := range respUpdate.Resp {
					h.Repo.UpdateDataIDChecksumFromServer(d.LocalID, "", d.CheckSum)
				}
			}

			//затем отправляем новые данные
			if len(req.Send.Data) == 0 {
				fmt.Println("Нет новых данных для отправки на сервер.")
			} else {
				respSend, err := h.Client.SendDataToServer(ctx, &req.Send)
				if err != nil {
					fmt.Printf("Возникла ошибка при отправке данных на сервер:\n%s", err.Error())
					return
				}

				fmt.Println("Новые данные отправлены на сервер.")
				for _, d := range respSend.Resp {
					h.Repo.UpdateDataIDChecksumFromServer(d.LocalID, d.Id, d.Checksum)
				}
			}

			//в конце скачиваем новые данные и обновления с сервера
			respGet, err := h.Client.GetDataFromServer(ctx, &req.Get)
			if err != nil {
				fmt.Printf("Возникла ошибка при скачивании данных с сервера:\n%s", err.Error())
				return
			}

			var dataToSave []model.UserData
			var dataToUpdate []model.UserData
			localID := h.Repo.GetNewLocalID()

			for _, d := range respGet.Data {
				s := model.UserData{
					ID:       d.DataID,
					Checksum: d.CheckSum,
					Title:    d.Title,
					Metadata: d.Metadata,
				}

				switch d.DataType {
				case LOGINTYPE:
					s.Data, err = encryption.Decode[model.LoginPassword](d.Userdata, *masterKey)
				case BINARYTYPE:
					s.Data, err = encryption.Decode[model.Binary](d.Userdata, *masterKey)
				case BANKCARDTYPE:
					s.Data, err = encryption.Decode[model.BankCard](d.Userdata, *masterKey)
				case TEXTTYPE:
					s.Data, err = encryption.Decode[model.Text](d.Userdata, *masterKey)
				}

				if err != nil {
					fmt.Printf("Возникла ошибка дешифровки данных после получения с сервера:\n%s", err.Error())
					return
				}

				if h.Repo.IsDataExistByID(d.DataID) {
					dataToUpdate = append(dataToUpdate, s)
				} else {
					s.LocalID = strconv.Itoa(localID)
					dataToSave = append(dataToSave, s)
					localID++
				}

			}

			h.Repo.UpdateDataByServerID(dataToUpdate)
			h.Repo.AddData(dataToSave)

			fmt.Println("Синхронизация завершена.")

		} else {
			fmt.Println("Для отправки данных нужно авторизоваться.")
		}
	}
}
