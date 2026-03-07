package models

import (
	"time"
)

type RefreshToken struct {
	ID        string `gorm:"primaryKey"`
	Token     string `gorm:"unique;index"`
	UserID    string `gorm:"index"`
	ExpiresAt time.Time
}

func (rt *RefreshToken) GetID() string {
	return rt.ID
}

func (rt *RefreshToken) GetToken() string {
	return rt.Token
}

func (rt *RefreshToken) GetUserID() string {
	return rt.UserID
}

func (rt *RefreshToken) GetExpiresAt() time.Time {
	return rt.ExpiresAt
}

func (rt *RefreshToken) IsExpired() bool {
	return time.Now().After(rt.ExpiresAt)
}

func (rt *RefreshToken) SetID(id string) {
	rt.ID = id
}

func (rt *RefreshToken) SetToken(token string) {
	rt.Token = token
}

func (rt *RefreshToken) SetUserID(userID string) {
	rt.UserID = userID
}

func (rt *RefreshToken) SetExpiresAt(expiresAt time.Time) {
	rt.ExpiresAt = expiresAt
}
