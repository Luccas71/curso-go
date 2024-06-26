package main

import (
	"database/sql"
	"fmt"

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
	product.Price = 1890.20
	err = updateProduct(db, product)
	if err != nil {
		panic(err)
	}
	// p, err := selectProduct(db, product.ID)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("Product: %v possui o valor de R$ %.2f\n", p.Name, p.Price)

	products, err := selectAllProducts(db)
	if err != nil {
		panic(err)
	}
	for _, p := range products {
		fmt.Printf("Product: %v, Price: %.2f\n", p.Name, p.Price)
	}

	err = removeProduct(db, product.ID)
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
	defer stmt.Close()

	//executando o stmt para trocar as ? pelos valores
	_, err = stmt.Exec(product.ID, product.Name, product.Price)
	if err != nil {
		panic(err)
	}
	return nil
}

func updateProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("update products set name = ?, price = ? where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(product.Name, product.Price, product.ID)
	if err != nil {
		return err
	}
	return nil
}

func selectProduct(db *sql.DB, id string) (*Product, error) {
	stmt, err := db.Prepare("select id, name, price from products where id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	//definir uma variavel para receber os dados da consulta
	var p Product
	//QueryRow => vai buscar apenas uma linha
	//Scan => vai substituir os dados da consulta nna variavel criada
	//passar o local na memoria para que os dados sejam alterados na variavel e nao em uma copia dela
	err = stmt.QueryRow(id).Scan(&p.ID, &p.Name, &p.Price)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func selectAllProducts(db *sql.DB) ([]Product, error) {
	//nao necessita de Prepare, pois nao será passado nenhum parametro sql
	rows, err := db.Query("select id, name, price from products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var products []Product
	//vai percorrer cada uma das linhas do resultado da consulta
	for rows.Next() {
		var p Product
		err = rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func removeProduct(db *sql.DB, id string) error {
	stmt, err := db.Prepare("delete from products where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}
