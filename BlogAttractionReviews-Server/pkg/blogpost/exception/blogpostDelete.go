package exception

type BlogPostDelete struct{}

func (e *BlogPostDelete) Error() string {
	return "blog post delete failed"
}
