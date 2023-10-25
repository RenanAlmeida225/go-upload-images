package schemas

import "time"

type Confirmation struct {
	ID        uint   `gorm:"primarykey"`
	Token     string `gorm:"unique"`
	UserID    uint
	User      User `gorm:"foreignKey:UserID"`
	CreatedAt time.Time
}
