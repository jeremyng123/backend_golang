package main

import (
	"log"

	fiber "github.com/gofiber/fiber/v2"
)

var app *fiber.App = fiber.New()

func main() {
	app.Get("/person/:id?", GetPerson)
	app.Post("/person", CreatePerson)
	app.Put("/person/:id", UpdatePerson)
	app.Delete("/person/:id", DeletePerson)
	err := app.Listen(port)
	if err != nil {
		log.Fatal("Server exited with error message: ", err)
	}
}
