package exception

type PasswordHashError struct{}

func (e *PasswordHashError) Error() string {
	return "password hash error"
}
