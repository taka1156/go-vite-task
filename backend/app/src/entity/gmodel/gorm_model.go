package gmodel

import "time"

type User struct {
	UserID    uint `gorm:"primary_key"`
	UserName  string
	UserIcon  *string
	Email     string
	Password  string
	RoleID    uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type Role struct {
	RoleID    uint `gorm:"primary_key"`
	RoleName  string
	RoleIcon  *string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
