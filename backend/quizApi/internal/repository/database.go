package repository

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"honnef.co/go/tools/config"
)

func NewDatabase(config *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DBUsername,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}

	// Auto Migrate Models
	err = db.AutoMigrate(
		&models.User{},
		&models.Admin{},
		&models.Participant{},
		&models.Quiz{},
		&models.Question{},
		&models.Option{},
	)
	if err != nil {
		return nil, err
	}

	return db, nil
}
