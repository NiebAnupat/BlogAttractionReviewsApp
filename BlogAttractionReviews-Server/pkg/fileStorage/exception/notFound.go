package exception

type NotFound struct {
	Name string
}

func (e *NotFound) Error() string {
	return "file not found: " + e.Name
}
