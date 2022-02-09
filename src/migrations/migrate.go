package migrations

import (
	"github.com/mailwilliams/auth/src/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewGORMConnection() (*gorm.DB, error) {

	return gorm.Open(mysql.Open("root:password@tcp(db:3306)/auth"), &gorm.Config{})
}

func AutoMigrate(db *gorm.DB) error {
	//	add list of models to migrate to the database here
	return db.AutoMigrate(models.User{})
}
