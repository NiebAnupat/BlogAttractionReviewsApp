package model

import "io"

type (
	ImageFileCreateReq struct {
		FileName    string    `json:"fileName"`
		FileType    string    `json:"fileType"`
		FileContent io.Reader `json:"fileContent"`
	}
)
