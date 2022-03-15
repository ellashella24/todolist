package util

import (
	"fmt"
	"log"

	"todolist-be/config"
	"todolist-be/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(config *config.AppConfig) *gorm.DB {
	connectionString := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",
		config.Username,
		config.Password,
		config.Address,
		config.DB_Port,
		config.Name,
	)

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		log.Println("failed to connect database :", err)
		panic(err)
	}

	InitMigrate(db)

	return db
}

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&entity.Todolist{})
}