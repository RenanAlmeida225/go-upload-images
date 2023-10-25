package schemas

import "time"

type Images struct {
	ID           uint `gorm:"primarykey"`
	Title        string
	Description  string
	Name         string
	OriginalName string
	MimeType     string
	Url          string
	UserId       uint
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
