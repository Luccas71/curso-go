package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	CategoryId int
	Category   Category
	//adc created_at, updated_at e deleted_at automaticos
	gorm.Model
}

type Category struct {
	ID   int `gorm:"primaryKey"`
	Name string
}

func main() {
	dns := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, Category{})

	//create category
	// category := Category{Name: "Eletronicos"}
	// db.Create(&category)

	//create product
	product := Product{
		Name:       "Keyboard",
		Price:      36.0,
		CategoryId: 1,
	}
	db.Create(&product)

	var products []Product
	db.Preload("Category").Find(&products)
	for _, product := range products {
		fmt.Println(product.Name, product.Category.Name)
	}
}
