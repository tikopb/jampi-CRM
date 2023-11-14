package model

import (
	"time"

	"github.com/google/uuid"
)

type Costumer struct {
	ID          int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string    `json:"name" gorm:"column:name;not null;index:idx_customer_name"`
	PhoneNumber string    `gorm:"phonenumber;not null;index:idx_customer_phonenumber" json:"phonenumber"`
	CreatedAt   time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	IsActive    bool      `gorm:"column:isactive;default:true" json:"IsActive"`
	UUID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();index:idx_customer_uuid"`
}

type CostumerRegister struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phonenumber"`
}
