package exception

type VerifyToken struct{}

func (e *VerifyToken) Error() string {
	return "verify token error"
}
