package main

import (
	"fmt"
	"log"

	"postgres/middleware" // <-- replace 'yourproject' with your actual module name
)

func main() {
	// Create database connection
	db := middleware.CreateConnection()
	defer db.Close()

	// Run a simple SQL query to verify connection
	var currentTime string
	err := db.QueryRow("SELECT NOW()").Scan(&currentTime)
	if err != nil {
		log.Fatal("Error executing test query:", err)
	}

	fmt.Println("✅ Connected to PostgreSQL successfully!")
	fmt.Println("Current database time:", currentTime)
}
