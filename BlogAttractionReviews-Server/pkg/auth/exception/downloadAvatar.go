package exception

type DownloadAvatar struct{}

func (e *DownloadAvatar) Error() string {
	return "Failed to download avatar"
}
