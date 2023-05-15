package handlers

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"gophkeeper/client/storage/model"
	"io"
	"os"
)

func (h Handlers) AddData(title *string, metadata *string, data any) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		h.Repo.AddData([]model.UserData{{
			LocalID:  h.Repo.GetNewLocalID(),
			Title:    *title,
			Metadata: *metadata,
			Data:     data,
		}})
		fmt.Println("данные успешно добавлены")
	}
}

func (h Handlers) EditData(localID *string, title *string, metadata *string, data any) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		h.Repo.UpdateDataByLocalID(model.UserData{
			LocalID:  *localID,
			Title:    *title,
			Metadata: *metadata,
			Checksum: "",
			Data:     data,
		})
		fmt.Println("данные успешно изменены")
	}
}

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

func (h Handlers) GetUserDataList() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		data := h.Repo.GetUserData()
		if len(data.Data) == 0 {
			fmt.Println("нет данных для отображения")
		} else {
			fmt.Println("список сохранённых данных")
			for _, v := range data.Data {
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

func (h Handlers) GetUserData(index *string) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		data := h.Repo.GetUserDataByLocalID(*index)
		fmt.Printf("Title: %s\nMetadata: %s\n\nData\n%s", data.Title, data.Metadata, data.Data)
	}
}
