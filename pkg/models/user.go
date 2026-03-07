package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID           string        `gorm:"primaryKey"`
	Name         string        `gorm:"not null"`
	Email        string        `gorm:"unique"`
	Password     []byte        `gorm:"not null"`
	RefreshToken *RefreshToken `gorm:"foreignKey:UserID"`
}

func (u *User) GetID() string {
	return u.ID
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetPassword() []byte {
	return u.Password
}

func (u *User) GetRefreshToken() *RefreshToken {
	return u.RefreshToken
}

func (u *User) SetID(id string) {
	u.ID = id
}

func (u *User) SetName(name string) {
	u.Name = name
}

func (u *User) SetEmail(email string) {
	u.Email = email
}

func (u *User) SetPassword(password []byte) {
	u.Password = password
}
