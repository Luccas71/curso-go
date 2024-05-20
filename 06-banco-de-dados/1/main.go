package main

import (
	"database/sql"

	//colocar o _ para nao travar o programa pois o mysql nao está sendo diretamente usado
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}
func main() {

	// estabelecendo conexão com db
	// db, err := sql.Open("db", "user:password@tcp(localhost:3306)/nomeDoBanco")
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	product := NewProduct("Notebook", 2000.00)
	err = insertProduct(db, product)
	if err != nil {
		panic(err)
	}
}

func insertProduct(db *sql.DB, product *Product) error {
	//preparando stmt para bloquear sql injection
	stmt, err := db.Prepare("insert into products(id, name, price) values (?, ?, ?)")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//executando o stmt para trocar as ? pelos valores
	_, err = stmt.Exec(product.ID, product.Name, product.Price)
	if err != nil {
		panic(err)
	}
	return nil
}
