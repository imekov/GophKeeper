package filesystem

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"os"

	"gophkeeper/client/storage/model"
)

// FS содержит имя файла и пользательские параметры вместе с данными.
type FS struct {
	filename    string
	UserSession model.UserSession
}

// NewFSStorage возвращает новый экземпляр для работы в файловом режиме.
func NewFSStorage(filename string) *FS {

	newFS := FS{
		filename: filename,
	}

	gob.Register(model.LoginPassword{})
	gob.Register(model.Text{})
	gob.Register(model.Binary{})
	gob.Register(model.BankCard{})

	err := newFS.ReadFile(&newFS.UserSession)
	if err != nil {
		if err = newFS.SaveFile(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	return &newFS
}

// OpenFile открывает файл.
func (s FS) OpenFile(flag int) *os.File {

	dataFile, err := os.OpenFile(s.filename, flag|os.O_CREATE, 0777)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return dataFile
}

// ReadFile читает данные из файла.
func (s FS) ReadFile(data *model.UserSession) error {

	dataFile := s.OpenFile(os.O_RDONLY)

	defer func(File *os.File) {
		err := File.Close()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}(dataFile)

	err := gob.NewDecoder(dataFile).Decode(data)
	if err != nil {
		s.UserSession = model.UserSession{
			Token:     "",
			DataArray: []model.UserData{},
		}
		return err
	}

	return nil
}

// SaveFile сохраняет данные в файл.
func (s FS) SaveFile() error {

	dataFile := s.OpenFile(os.O_WRONLY)
	defer dataFile.Close()

	writer := bufio.NewWriter(dataFile)

	err := gob.NewEncoder(writer).Encode(&s.UserSession)
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = writer.Flush()
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil

}
