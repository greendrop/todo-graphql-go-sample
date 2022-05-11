package entity

import (
	"time"
)

type User struct {
	Id          int64 // `gorm:"primary_key AUTO_INCREMENT"`
	name        *string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
