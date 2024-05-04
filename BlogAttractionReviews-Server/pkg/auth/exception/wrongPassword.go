package exception

type WrongPassword struct{}

func (e *WrongPassword) Error() string {
	return "wrong password"
}
