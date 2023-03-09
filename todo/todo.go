package todo

import (
	"github.com/gofiber/fiber"
)

func GetTodos(c *fiber.Ctx) {
	c.JSON(todos)
}

func CreateTodo(c *fiber.Ctx) {
	todo := new(Todo)

	if err := c.BodyParser(todo); err != nil {
		c.Status(503).Send(err)
		return
	}

	todos = append(todos, *todo)
	c.JSON(todo)
}
