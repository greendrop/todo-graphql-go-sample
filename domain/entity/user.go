package entity

import (
	"time"
)

type User struct {
	Id        int64 // `gorm:"primary_key AUTO_INCREMENT"`
	Name      *string
	CreatedAt time.Time
	UpdatedAt time.Time
}
