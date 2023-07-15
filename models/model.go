package models

import (
	// ThirdParty libs
	"gorm.io/gorm"
)

type User struct {
	ID       uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name     string  `gorm:"not null" json:"name"`
	Email    string  `gorm:"unique;not null" json:"email"`
	Password string  `gorm:"not null" json:"password"`
	Session  Session `gorm:"foreignKey:Email;references:Email" json:"session"`
	gorm.Model
}

type Note struct {
	ID    uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Email string `gorm:"not null" json:"email"`
	Note  string `gorm:"not null" json:"note"`
	gorm.Model
}

type Session struct {
	Email     string `gorm:"primaryKey;not null" json:"email"`
	SessionID string `gorm:"unique;not null" json:"session_id"`
}

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User) TableName() string {
	return "users"
}

func (u *Session) TableName() string {
	return "sessions"
}

func (u *Note) TableName() string {
	return "notes"
}
