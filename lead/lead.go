package lead

import (
	"github.com/Aakashraz/crm-with-golang-fiber/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
)

type Lead struct {
	gorm.Model
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   int    `json:"phone"`
}

func GetLeads(c *fiber.Ctx) {
	db := database.Db
	var leads []Lead
	db.Find(&leads)
	err := c.JSON(leads)
	if err != nil {
		log.Printf("error while serializing and sending JSON response: %v", err)
	}
}

func GetLead(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.Db
	var ld Lead
	//Executes a database query to find a lead with the specified ID (value of the "id" parameter) and stores the result in the ld variable.
	db.Find(&ld, id)
	//Uses the c.JSON method provided by the Fiber framework
	//to serialize the ld (Lead) struct into JSON format and send it as the response.
	err := c.JSON(ld)
	if err != nil {
		log.Printf("error while serializing and sending JSON response: %v", err)
		return
	}
}

func NewLead(c *fiber.Ctx) {
	db := database.Db
	//The new function is used to allocate memory for a new	  zero-initialized value of the specified type.
	lead := new(Lead)
	if err := c.BodyParser(lead); err != nil {
		c.Status(503).Send(err)
		return
	}
	db.CreateTable(&lead)
	err := c.JSON(lead)
	if err != nil {
		log.Printf("error while serializing and sending JSON response: %v", err)
	}
}

func DeleteLead(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.Db

	var ld Lead
	db.First(&ld, id)
	if ld.Name == "" {
		c.Status(500).Send("No lead found with the ID.")
		return
	}
	db.Delete(&ld)
	c.Send("lead successfully Deleted!!!")
}
