package db

import (
	"fmt"

	"github.com/vkhoa145/go-chat-next-ts/app/models"
	"github.com/vkhoa145/go-chat-next-ts/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Ho_Chi_Minh",
		cfg.DB.Host,
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Name,
		cfg.DB.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	db.AutoMigrate(
		&models.User{},
	)

	if err != nil {
		panic(err.Error())
	}

	return db
}
