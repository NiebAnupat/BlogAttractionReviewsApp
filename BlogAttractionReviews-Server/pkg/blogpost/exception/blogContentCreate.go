package exception

type BlogContentCreate struct{}

func (e BlogContentCreate) Error() string {
	return "Creating blog content failed"
}
