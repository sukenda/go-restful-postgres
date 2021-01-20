package config

import (
	"github.com/sukenda/go-restful-postgres/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB
var err error

func NewPostgresDatabase(configuration Config) *gorm.DB {
	dbHost := configuration.Get("DATABASE_HOST")
	dbPort := configuration.Get("DATABASE_PORT")
	dbUser := configuration.Get("DATABASE_USERNAME")
	dbPass := configuration.Get("DATABASE_PASSWORD")
	dbName := configuration.Get("DATABASE_NAME")

	dsn := "host=" + dbHost + " port=" + dbPort + " user=" + dbUser + " password=" + dbPass + " dbname=" + dbName + " sslmode=disable"
	database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&entity.User{})
	database.AutoMigrate(&entity.Department{})
	database.AutoMigrate(&entity.Employee{})
	database.AutoMigrate(&entity.Project{})
	database.AutoMigrate(&entity.Team{})

	return database
}

func GetDatabase() *gorm.DB {
	return database
}
