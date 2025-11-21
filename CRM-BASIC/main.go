package main

import (
	"crm/database"
	"crm/lead"
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.CreateLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)

}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "leads.db")
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Database connection established")
	database.DBConn.AutoMigrate(&lead.lead{})
	fmt.Println("Database migrated")

}
func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	app.Listen(":3000")
	defer database.DBConn.Close()
}
