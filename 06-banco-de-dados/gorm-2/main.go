package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Price        float64
	CategoryId   int
	Category     Category
	SerialNumber SerialNumber
	//adc created_at, updated_at e deleted_at automaticos
	gorm.Model
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    string
	ProductId int
}

type Category struct {
	ID   int `gorm:"primaryKey"`
	Name string
	//pode ter varios produtos
	Products []Product
}

func main() {
	dns := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, Category{}, SerialNumber{})

	//create category
	// category := Category{Name: "Cozinha"}
	// db.Create(&category)

	// //create product
	// product := Product{
	// 	Name:       "Panela",
	// 	Price:      32.60,
	// 	CategoryId: 1,
	// }
	// db.Create(&product)

	// //create serialNumber
	// db.Create(&SerialNumber{
	// 	Number:    "123",
	// 	ProductId: 1,
	// })

	var categories []Category
	err = db.Model(&Category{}).Preload("Products.SerialNumber").Find(&categories).Error
	if err != nil {
		panic(err)
	}
	for _, category := range categories {
		fmt.Println(category.Name, ":")
		for _, product := range category.Products {
			fmt.Println("- ", product.Name, " - Serial Number: ", product.SerialNumber.Number)
		}
	}
}
