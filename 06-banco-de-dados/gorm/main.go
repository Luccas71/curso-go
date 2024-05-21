package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
	//adc created_at, updated_at e deleted_at automaticos
	gorm.Model
}

func main() {
	dns := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=true&loc=Local"
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

	//update
	var p Product
	db.First(&p, 1)
	// p.Name = "bola de capotão"
	// db.Save(&p)
	// fmt.Println(p.Name)

	// // delete
	db.Delete(&p)

	// SOFT DELETE => as infos sao deletadas mas permanecem no db
}
