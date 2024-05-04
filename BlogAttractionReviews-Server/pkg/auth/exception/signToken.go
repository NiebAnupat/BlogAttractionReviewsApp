package exception

type SignToken struct{}

func (e *SignToken) Error() string {
	return "sign token error"
}
