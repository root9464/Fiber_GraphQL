package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
type Data struct {
	DB *gorm.DB
}

var Db Data
func ConnectToDB() (Data, error) {
	dsn := "блять вот сюда ссылку на бд"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("не удалось подключиться к базе данных")
	}
	Db = Data{
		DB: db,
	}
	return Db, nil
}
