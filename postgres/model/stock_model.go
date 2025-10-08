package model

import (
	"fmt"
	"log"
	"postgres/middleware"
)

type Stock struct {
	StockID int    `json:"id"`
	Name    string `json:"name"`
	Price   int    `json:"price"`
	Company string `json:"company"`
}

// // Returns SQL to create table
// func CreateTable() string {
// 	return `CREATE TABLE IF NOT EXISTS stock (
// 		stock_id SERIAL PRIMARY KEY,
// 		name VARCHAR(100) NOT NULL,
// 		price INT NOT NULL,
// 		company VARCHAR(100) NOT NULL
// 	);`
// }

// Insert a stock record into the database
func CreateStock(s Stock) error {
	db := middleware.CreateConnection()
	defer db.Close()
	sqlStatement := `
	INSERT INTO stock (name, price, company)
	VALUES($1, $2, $3)
	RETURNING id;
	`
	id := 0
	err := db.QueryRow(sqlStatement, s.Name, s.Price, s.Company).Scan(&id)
	if err != nil {
		panic(err)
	}

	fmt.Println("New record ID is:", id)
	return nil
}
func GetStocks() ([]Stock, error) {
	db := middleware.CreateConnection()
	defer db.Close()
	var stocks []Stock
	sqlStatement := `SELECT * FROM stock`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var s Stock
		err := rows.Scan(&s.StockID, &s.Name, &s.Price, &s.Company)
		if err != nil {
			log.Fatal(err)
		}
		stocks = append(stocks, s)
	}
	for i := range stocks {
		fmt.Println(stocks[i])
	}
	return stocks, nil
}
