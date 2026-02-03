package model

import "errors"

var (
	ErrNotFound      = errors.New("not found")
	ErrAlreadyExists = errors.New("already exists")
)

type Shortening struct {
	Identifier  string `json:"identifier"`
	OriginalURL string `[jxejson:"original_url"`
	Visits      int    `json:"visits"`
}
