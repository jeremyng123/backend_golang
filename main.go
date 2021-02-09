package main

import fiber "github.com/gofiber/fiber/v2"

var app *fiber.App = fiber.New()

func main() {
	app.Get("/person/:id?", GetPerson)
	app.Post("/person", CreatePerson)
	app.Put("/person/:id", UpdatePerson)
	app.Delete("/person/:id", DeletePerson)
	app.Listen(port)
}
