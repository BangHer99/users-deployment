package migration

import (
	Users "be12/deploy/features/users/data"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&Users.User{})
}
