package schemas

import (
	"time"
)

type User struct {
	ID        uint `gorm:"primarykey"`
	Username  string
	Email     string `gorm:"unique"`
	Passwprd  string
	IsEnable  string `gorm:"default:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
