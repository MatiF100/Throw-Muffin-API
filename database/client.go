package database

import (
	"log"

	"github.com/MatiF100/Throw-Muffin-API/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var Instance *gorm.DB
var dbError error

func Connect(connectionString string) {
	Instance, dbError = gorm.Open(postgres.Open(connectionString), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if dbError != nil {
		log.Fatal(dbError)
		panic("failed to connect database")
	}
	log.Println("Database connected")
}

func Migrate() {
	Instance.AutoMigrate(&models.User{}, &models.BodyPart{}, &models.Excercise{})
	log.Println("Database migrated")
}
