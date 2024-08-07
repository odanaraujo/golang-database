package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"log"
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
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Println("não foi possível acessar o banco de dados")
		panic(err)
	}

	/*product := NewProduct("Camisa", 80.00)
	product.ID = "82b3baf1-97a6-4496-a2af-132a9ea41b6e"
	if err := updateProduct(db, product); err != nil {
		fmt.Println("erro ao inserir os dados")
		panic(err)
	}*/
	/*prod, err := getProduct(db, "82b3baf1-97a6-4496-a2af-132a9ea41b6e")
	if err != nil {
		panic(err)
	}*/
	/*product := NewProduct("Sapato", 280.00)
	if err := insertProducts(db, product); err != nil {
		fmt.Println("erro ao inserir os dados")
		panic(err)
	}*/

	if err = deleteProduct(db, "82b3baf1-97a6-4496-a2af-132a9ea41b6e"); err != nil {
		panic(err)
	}

	prods, err := getProducts(db)
	if err != nil {
		fmt.Println("error")
		panic(err)
	}
	fmt.Println(prods)
}

func insertProducts(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("insert into products (id, name, price) values (?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	if _, err = stmt.Exec(product.ID, product.Name, product.Price); err != nil {
		return err
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

func getProduct(db *sql.DB, id string) (*Product, error) {
	stmt, err := db.Prepare("select * from products where id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var product Product

	//podemos mudar para stmt.QueryRowContext - aqui ele espera um contexto e os argumentos
	if err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price); err != nil {
		return nil, err
	}
	return &product, nil
}

func getProducts(db *sql.DB) (*[]Product, error) {
	rows, err := db.Query("select id, name, price from products")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var products []Product

	for rows.Next() {
		var product Product
		if err = rows.Scan(&product.ID, &product.Name, &product.Price); err != nil {
			continue
		}
		products = append(products, product)
	}
	return &products, nil
}

func deleteProduct(db *sql.DB, id string) error {
	stmt, err := db.Prepare("delete from products where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err = stmt.Exec(id); err != nil {
		return err
	}
	return nil
}
