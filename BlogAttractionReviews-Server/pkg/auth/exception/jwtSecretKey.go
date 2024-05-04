package exception

type JWTSecretKey struct{}

func (e *JWTSecretKey) Error() string {
	return "Required JWT secret key"
}
