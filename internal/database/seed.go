package database

import (
	"JampiCrm/internal/model"

	"gorm.io/gorm"
)

func seedDB(db *gorm.DB) {
	db.AutoMigrate(
		&model.User{},
	)
}
