package schemas

import (
	"time"
)

type User struct {
	ID        uint `gorm:"primarykey"`
	Username  string
	Email     string `gorm:"unique"`
	Password  string
	IsEnable  bool `gorm:"default:false"`
	Images    []Images
	CreatedAt time.Time
	UpdatedAt time.Time
}
