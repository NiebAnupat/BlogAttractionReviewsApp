package model

import (
	"mime/multipart"
	"time"
)

const (
	BlogContentText int = iota
	BlogContentImage
)

type (
	BlogPostCreateFormReq struct {
		Title       string                `json:"title"`
		Description string                `json:"description" `
		Thumbnail   *multipart.FileHeader `json:"thumbnail,omitempty"`
	}

	BlogPostCreateReq struct {
		Title       string
		Description string
		Thumbnail   string
		AuthorID    string
	}

	BlogContentCreateReq struct {
		BlogID string
		Order  int
		// Type 1 = Image, 0 = Text
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
