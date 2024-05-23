package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
	//pode ter varias categorias
	Categories []Category `gorm:"many2many:products_categories;"`
	//adc created_at, updated_at e deleted_at automaticos
	gorm.Model
}

type Category struct {
	ID   int `gorm:"primaryKey"`
	Name string
	//pode ter varios produtos
	Products []Product `gorm:"many2many:products_categories;"`
}

func main() {
	dns := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, Category{})

	//create category
	category1 := Category{Name: "Cozinha"}
	db.Create(&category1)

	category2 := Category{Name: "Eletronicos"}
	db.Create(&category2)

	// //create product
	product := Product{
		Name:       "Panela",
		Price:      32.60,
		Categories: []Category{category1, category2},
	}
	db.Create(&product)

	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Find(&categories).Error
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		fmt.Println(category.Name, ":")
		for _, product := range category.Products {
			fmt.Println(product.Name)
		}
	}
}
