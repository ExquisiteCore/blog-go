package model

import (
	"backend/config"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDb() {
	dbConfig := config.GlobalConfig.Database

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.DbHost, dbConfig.DbPort, dbConfig.DbUser, dbConfig.DbPassWord, dbConfig.DbName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{}, &Post{}, &Category{})

}
