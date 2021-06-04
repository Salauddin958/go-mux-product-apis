package main

import (
	"database/sql"
	"fmt"
)

type product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (p *product) getProduct(db *sql.DB) error {
	// rows, err := db.Query("SELECT name, price FROM products WHERE id=?", p.ID)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer rows.Close()
	// for rows.Next() {
	// 	err := rows.Scan(&p.Name, &p.Price)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }
	// return err
	return db.QueryRow("SELECT name, price FROM products WHERE id=?",
		p.ID).Scan(&p.Name, &p.Price)
}

func (p *product) updateProduct(db *sql.DB) error {
	_, err :=
		db.Exec("UPDATE products SET name=?, price=? WHERE id=?",
			p.Name, p.Price, p.ID)

	return err
}

func (p *product) deleteProduct(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM products WHERE id=?", p.ID)
	return err
}

func (p *product) createProduct(db *sql.DB) error {
	query := "INSERT INTO products(name, price) VALUES(?, ?)"
	data, err := db.Exec(query, p.Name, p.Price)
	id, err := data.LastInsertId()
	p.ID = (int)(id)
	if err != nil {
		fmt.Println("err in create : ", err)
		return err
	}
	return nil
}

func getProducts(db *sql.DB, start, count int) ([]product, error) {
	rows, err := db.Query(
		"SELECT id, name,  price FROM products LIMIT ? OFFSET ?",
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	products := []product{}

	for rows.Next() {
		var p product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}
