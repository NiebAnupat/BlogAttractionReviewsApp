package exception

import "fmt"

type UserNotFound struct {
	Username string
}

func (e UserNotFound) Error() string {
	return fmt.Sprintf("user with username %s not found", e.Username)
}
