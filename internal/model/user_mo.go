package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        string    `gorm:"primaryKey;autoIncrement" json:"-"`
	Username  string    `gorm:"unique" json:"username"`
	Hash      string    `json:"-"`
	FullName  string    `gorm:"column:full_name" json:"full_name"`
	CreatedAt time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	IsActive  bool      `gorm:"column:isactive" json:"isactive"`
	UUID      uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();index:idx_user_uuid"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	FullName string `json:"fullname"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserSession struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type UserSessionRespond struct {
	UserInformation User        `json:"user_information"`
	UserSession     UserSession `json:"user_session"`
}

type RefreshSession struct {
	RefreshToken string `json:"refresh_token"`
}

type AccessToken struct {
	AccessToken string `json:"access_token"`
}
