package lead

import (
	"crm/database"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type lead struct {
	gorm.Model

	Name    string
	Company string
	Email   string
	Phone   int
}

func GetLeads(c *fiber.Ctx) {
	db := database.DBConn
	var leads []lead
	db.Find(&leads)
	c.JSON(leads)
}
func GetLead(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var lead lead
	db.Find(&lead, id)
	c.JSON(lead)
}
func CreateLead(c *fiber.Ctx) {
	db := database.DBConn
	lead := new(lead)
	if err := c.BodyParser(lead); err != nil {
		c.Status(503).Send(err)
	}
	db.Create(&lead)
	c.JSON(lead)

}

func DeleteLead(c *fiber.Ctx) {
	id := c.Params("id")
	var lead lead
	db := database.DBConn
	db.First(&lead, id)
	if lead.Name == "" {
		c.Status(404).Send("No lead found with given ID")
		return
	}
	db.Delete(&lead)
	c.Send("Lead deleted successfully")

}
