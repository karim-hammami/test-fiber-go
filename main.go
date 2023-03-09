package main

import (
	"github.com/gofiber/fiber"
	"github.com/karim-hammami/test-fiber-go/todo"
)

func helloWorld(c *fiber.Ctx) {
	c.Send("Hello World!")
}

func setupRoutes(app *fiber.App) {
	app.Get("/", helloWorld)

	app.Get("/todos", todo.GetTodos)
	app.Post("/todos", todo.CreateTodo)
}

func main() {
	app := fiber.New()

	setupRoutes(app)

	app.Listen(3000)
}
