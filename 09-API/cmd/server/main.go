package main

import (
	"net/http"

	"github.com/Luccas71/curso-go/09-API/configs"
	"github.com/Luccas71/curso-go/09-API/internal/entity"
	"github.com/Luccas71/curso-go/09-API/internal/infra/database"
	"github.com/Luccas71/curso-go/09-API/internal/infra/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	_, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.User{}, &entity.Product{})

	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	http.HandleFunc("/products", productHandler.CreateProduct)
	http.ListenAndServe(":8000", nil)
}
