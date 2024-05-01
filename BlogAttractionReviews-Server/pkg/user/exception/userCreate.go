package exception

import "fmt"

type UserCreate struct {
	ID string
}

func (e UserCreate) Error() string {
	return fmt.Sprintf("Creating user id %v failed", e.ID)
}
