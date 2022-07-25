package db

import (
	"log"

	"github.com/mahdi-asadzadeh/go-grpc-product-microservice/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(&models.ProductImage{})
	db.AutoMigrate(&models.Product{})
}

func Init(url string) Handler {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	migrate(db)

	return Handler{db}
}
