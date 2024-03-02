package main

import (
	"fmt"
	"github.com/Aakashraz/crm-with-golang-fiber/database"
	"github.com/Aakashraz/crm-with-golang-fiber/lead"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"log"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLeads)
}

func initDatabase() {
	var err error
	database.Db, err = gorm.Open("sqlite", "leads.db")
	if err != nil {
		log.Printf("error while datatbse connection:%s", err)
	}
	fmt.Println("Database Connection Established.")
	database.Db.AutoMigrate(&lead.Lead{})
	fmt.Println("Database Migrated.")
}

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	err := app.Listen(3000)
	if err != nil {
		return
	}
	defer database.Db.Close()
}
