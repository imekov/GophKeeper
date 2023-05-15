package filesystem

func (s FS) IsUserAuthorized() bool {
	return s.Data.Token != ""
}

func (s FS) UpdateToken(newToken string) {
	s.Data.Token = newToken
	s.SaveFile()
}
