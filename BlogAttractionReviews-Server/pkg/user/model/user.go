package model

type (
	UserCreateReq struct {
		Username string
		Email    string
		Password string
		Avatar   string
	}

	User struct {
		ID       string
		Username string
		Email    string
		Avatar   string
	}
)
