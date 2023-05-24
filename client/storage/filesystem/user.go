package filesystem

// IsUserAuthorized проверяет авторизирован ли текущий пользователь.
func (s FS) IsUserAuthorized() bool {
	return s.UserSession.Token != ""
}

// UpdateToken обновляет пользовательский токен.
func (s FS) UpdateToken(newToken string) {
	s.UserSession.Token = newToken
	s.SaveFile()
}
