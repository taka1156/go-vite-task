package db_model

import "time"

type User struct {
	UserId    uint `gorm:"primaryKey"`
	UserName  string
	UserIcon  *string
	Email     string
	Password  [16]byte
	RoleId    *uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
