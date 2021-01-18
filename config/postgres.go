package config

import (
	"github.com/sukenda/go-restful-postgre/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDatabase(configuration Config) *gorm.DB {
	dbHost := configuration.Get("DATABASE_HOST")
	dbPort := configuration.Get("DATABASE_PORT")
	dbUser := configuration.Get("DATABASE_USERNAME")
	dbPass := configuration.Get("DATABASE_PASSWORD")
	dbName := configuration.Get("DATABASE_NAME")

	dsn := "host=" + dbHost + " port=" + dbPort + " user=" + dbUser + " password=" + dbPass + " dbname=" + dbName + " sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&entity.User{})

	return database
}
