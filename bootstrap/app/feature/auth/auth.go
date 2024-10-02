package auth

type Auth struct {
	UserID   uint
	Email    string
	LoggedIn bool
}

func (auth Auth) Check() bool {
	return auth.LoggedIn
}
