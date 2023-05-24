package model

import "fmt"

// UserSession содержит токен и массив пользовательских данных.
type UserSession struct {
	Token     string
	DataArray []UserData
}

// UserData содержит пользовательские данные.
type UserData struct {
	LocalID  string
	ID       string
	Title    string
	Metadata string
	Checksum string
	Data     any
}

// LoginPassword содержит логин и пароль.
type LoginPassword struct {
	Login    string
	Password string
}

func (p LoginPassword) String() string {
	return fmt.Sprintf("Login:%s\nPassword:%s", p.Login, p.Password)
}

// Text содержит текстовые данные.
type Text struct {
	TextData string
}

func (p Text) String() string {
	return fmt.Sprintf("Textdata:%s", p.TextData)
}

// Binary содержит бинарные пользовательские данные.
type Binary struct {
	Path       string
	BinaryData []byte
}

func (p Binary) String() string {
	return fmt.Sprintf("Path:%s\nBinaryData:%s", p.Path, p.BinaryData)
}

// BankCard содержит банковские данные.
type BankCard struct {
	Number  string
	Owner   string
	ExpDate string
}

func (p BankCard) String() string {
	return fmt.Sprintf("Number:%s\nOwner:%s\nExpDate:%s", p.Number, p.Owner, p.ExpDate)
}
