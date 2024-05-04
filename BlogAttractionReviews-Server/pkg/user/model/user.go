package model

type (
	UserCreateReq struct {
		Username string
		Password string
		Avatar   string
	}

	User struct {
		ID       string
		Username string
		Avatar   string
	}
)
