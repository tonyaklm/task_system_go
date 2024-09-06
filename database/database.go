package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"task_system_go/config"
)

var Database *gorm.DB

func InitDatabase() error {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.Cfg.Database.Host,
		config.Cfg.Database.Username,
		config.Cfg.Database.Password,
		config.Cfg.Database.DBName,
		config.Cfg.Database.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	Database = db
	return nil
}
