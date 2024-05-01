package exception

type BlogPostNotFound struct {
	ID string
}

func (e *BlogPostNotFound) Error() string {
	return "Blog post not found with ID: " + e.ID
}
