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
}

func main() {
	dns := "root:root@tcp(localhost:3306)/goexpert"
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{})

	//criando um produto
	// db.Create(&Product{
	// 	Name:  "bola",
	// 	Price: 15.0,
	// })

	//criando varios produtos de uma vez
	// products := []Product{
	// 	{Name: "raquete", Price: 12.5},
	// 	{Name: "peteca", Price: 2.5},
	// 	{Name: "baldinho", Price: 10.99},
	// }

	// db.Create(&products)

	//select one
	// var product Product
	// db.First(&product, 2)
	// fmt.Println(product)

	//select por atributo
	// db.First(&product, "name = ?", "peteca")
	// fmt.Println(product)

	//select all
	// var products []Product
	// db.Find(&products)
	// for _, p := range products {
	// 	fmt.Println(p)
	// }

	//where
	// var products []Product
	// db.Where("price < ?", 10).Find(&products)
	// for _, p := range products {
	// 	fmt.Println(p)
	// }

	//like
	// var products []Product
	// db.Where("name LIKE ?", "%e%").Find(&products)
	// for _, p := range products {
	// 	fmt.Println(p)
	// }
}
