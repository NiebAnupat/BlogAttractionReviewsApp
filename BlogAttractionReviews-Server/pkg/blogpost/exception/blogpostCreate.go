package exception

type BlogPostCreate struct{}

func (e BlogPostCreate) Error() string {
	return "Creating blog post failed"
}
