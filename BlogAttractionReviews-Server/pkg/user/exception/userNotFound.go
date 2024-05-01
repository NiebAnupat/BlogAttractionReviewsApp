package exception

import "fmt"

type UserNotFound struct {
	ID string
}

func (e UserNotFound) Error() string {
	return fmt.Sprintf("user with id %s not found", e.ID)
}
