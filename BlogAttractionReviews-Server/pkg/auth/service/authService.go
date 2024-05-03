package service

type AuthService interface {
	Login(username, password string) (string, error)
	Logout(token string) error

	VerifyToken(token string) (string, error)

	RefreshToken(token string) (string, error)

	Register(username, email, password, avatar string) error
}
