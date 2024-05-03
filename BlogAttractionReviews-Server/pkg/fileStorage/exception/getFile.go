package exception

type GetFileFailed struct {
	Name string
}

func (e *GetFileFailed) Error() string {
	return "get file failed: " + e.Name
}
