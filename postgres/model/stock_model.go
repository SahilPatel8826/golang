package model

import (
	"fmt"
	"log"
	"postgres/middleware"
)

type Stock struct {
	StockID int     `json:"id"`
	Name    *string `json:"name"`
	Price   *int    `json:"price"`
	Company *string `json:"company"`
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
	RETURNING stock_id;
	`
	stock_id := 0
	err := db.QueryRow(sqlStatement, s.Name, s.Price, s.Company).Scan(&stock_id)
	if err != nil {
		panic(err)
	}

	fmt.Println("New record ID is:", stock_id)
	return nil
}
func GetAllStocks() ([]Stock, error) {
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
func GetStocks(id string) (Stock, error) {
	db := middleware.CreateConnection()
	defer db.Close()

	var s Stock
	sqlStatement := `SELECT * FROM stock WHERE stock_id = $1`
	row := db.QueryRow(sqlStatement, id)

	err := row.Scan(&s.StockID, &s.Name, &s.Price, &s.Company)
	if err != nil {
		return s, err
	}

	return s, nil
}

func DeleteStock(id string) {

	db := middleware.CreateConnection()
	defer db.Close()

	sqlStatement := `DELETE FROM stock WHERE stock_id = $1;`
	result, err := db.Exec(sqlStatement, id)
	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, _ := result.RowsAffected()
	fmt.Println("Deleted rows:", rowsAffected)

}
func UpdateStock(id string, s Stock) (Stock, error) {
	db := middleware.CreateConnection()
	defer db.Close()

	// First, fetch the current stock
	var current Stock
	err := db.QueryRow(`SELECT stock_id, name, price, company FROM stock WHERE stock_id=$1`, id).
		Scan(&current.StockID, &current.Name, &current.Price, &current.Company)
	if err != nil {
		return Stock{}, fmt.Errorf("stock not found: %v", err)
	}

	// Only update fields that are non-nil
	if s.Name != nil {
		current.Name = s.Name
	}
	if s.Price != nil {
		current.Price = s.Price
	}
	if s.Company != nil {
		current.Company = s.Company
	}

	// Update in DB
	sqlStatement := `
		UPDATE stock
		SET name=$1, price=$2, company=$3
		WHERE stock_id=$4
		RETURNING stock_id, name, price, company;
	`

	err = db.QueryRow(sqlStatement, current.Name, current.Price, current.Company, id).
		Scan(&current.StockID, &current.Name, &current.Price, &current.Company)
	if err != nil {
		return Stock{}, fmt.Errorf("update failed: %v", err)
	}

	return current, nil
}
