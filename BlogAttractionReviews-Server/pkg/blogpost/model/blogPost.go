package model

type (
	BlogPostCreateReq struct {
		Title    string
		Content  []BlogContentCreateReq
		AuthorID string
	}

	BlogContentCreateReq struct {
		Order    int
		Type     int
		Text     string
		ImageURL string
	}

	BlogPost struct {
		ID        string
		Title     string
		CreateAt  string
		Contens   []BlogContent
		AuthorID  string
		Likes     int
		Favorites int
	}

	BlogContent struct {
		ID       string
		Order    int
		Type     int
		Text     string
		ImageURL string
	}
)
