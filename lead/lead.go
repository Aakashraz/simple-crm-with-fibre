package lead

import (
	"github.com/Aakashraz/crm-with-golang-fiber/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"log"
)

type Lead struct {
	gorm.Model
	Name    string
	Company string
	Email   string
	Phone   int
}

func GetLeads(c *fiber.Ctx) {
	//...
}

func GetLead(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.Db
	var ld Lead
	//Executes a database query to find a lead with the specified ID (value of the "id" parameter) and stores the result in the ld variable.
	db.Find(&ld, id)
	//Uses the c.JSON method provided by the Fiber framework to serialize the ld (Lead) struct into JSON format and send it as the response.
	err := c.JSON(ld)
	if err != nil {
		log.Printf("error while reading json: %v", err)
		return
	}
}
