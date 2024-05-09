package model

import (
	"mime/multipart"
	"time"
)

const (
	BlogContentText ContentType = iota + 1
	BlogContentImage
)

type (
	ContentType           int
	BlogPostCreateFormReq struct {
		Title       string                `json:"title"`
		Description string                `json:"description" `
		Thumbnail   *multipart.FileHeader `json:"thumbnail,omitempty"`
	}

	BlogContentCreateFormReq struct {
		Order int         `json:"order" `
		Type  ContentType `json:"type"`
		Value interface{} `json:"value"`
	}

	BlogPostCreateReq struct {
		Title       string
		Description string
		Thumbnail   string
		AuthorID    string
	}

	BlogContentCreateReq struct {
		Order int
		Type  int
		Value string
	}

	BlogPost struct {
		ID          string
		Title       string
		Description string
		Thumbnail   string
		CreateAt    time.Time
		Contents    []*BlogContent
		AuthorID    string
		Likes       int
		Favorites   int
	}

	BlogContent struct {
		ID    string
		Order int
		Type  int
		Value string
	}
)
