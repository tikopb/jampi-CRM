package model

import (
	"time"

	"github.com/google/uuid"
)

type Checkup struct {
	ID          int       `gorm:"primaryKey;autoIncrement" json:"id"`
	CreatedAt   time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	Tensi       int       `gorm:"column:tensi" json:"tensi"`
	GulaDarah   int       `gorm:"column:guladarah" json:"guladarah"`
	Cholestrol  int       `gorm:"column:cholestrol" json:"cholestrol"`
	AsamUrat    int       `gorm:"column:asamurat" json:"asamurat"`
	Customer_id int       `gorm:"column:customer_id;not null;index:idx_checkup_customer_id" json:"created_by"`
	Costumer    Costumer  `gorm:"foreignKey:customer_id"`
	UUID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();index:idx_checkup_uuid"`
}
