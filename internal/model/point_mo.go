package model

import "github.com/google/uuid"

type Point struct {
	ID          int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Customer_id int       `gorm:"column:customer_id;index:idx_point_customer_id" json:"created_by"`
	Costumer    Costumer  `gorm:"foreignKey:customer_id"`
	UUID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();index:idx_point_uuid"`
}
