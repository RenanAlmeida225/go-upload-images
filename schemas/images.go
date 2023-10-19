package schemas

import "gorm.io/gorm"

type Images struct {
	gorm.Model
	Title        string
	Description  string
	Name         string
	OriginalName string
	MimeType     string
	Url          string
}
