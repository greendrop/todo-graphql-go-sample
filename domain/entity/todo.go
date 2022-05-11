package entity

import (
	"time"
)

type Todo struct {
	Id        int64 // `gorm:"primary_key AUTO_INCREMENT"`
	UserId    int64
	Title     *string
	Done      bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
