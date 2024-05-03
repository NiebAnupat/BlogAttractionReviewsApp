package exception

type UploadFailed struct {
	Name string
}

func (e *UploadFailed) Error() string {
	return "upload failed: " + e.Name
}
