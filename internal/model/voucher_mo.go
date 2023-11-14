package model

import (
	"time"

	"github.com/google/uuid"
)

type Voucher struct {
	ID          int       `gorm:"primaryKey;autoIncrement" json:"id"`
	ExpDate     time.Time `gorm:"column:expdate;not null;index:idex_voucher_expdate" json:"exp_date"`
	CreatedAt   time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP;index:idx_voucher_created_at;not null" json:"created_at"`
	Code        string    `gorm:"column:code" json:"code"`
	Used        bool      `gorm:"column:used;default:false" json:"used"`
	Customer_id int       `gorm:"column:customer_id;not null;index:idx_voucher_customer_id" json:"created_by"`
	Costumer    Costumer  `gorm:"foreignKey:customer_id"`
	UUID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();index:idx_voucher_uuid"`
}
