package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	ID        string `json:"id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

func main(){
	fmt.Println("hello world")

	app := fiber.New()

	todos := []Todo{}
	
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("Hello, World!")
	})

	app.Post("/todo", func(c *fiber.Ctx) error {
		todo := new(Todo)

		fmt.Println("Received request to create todo")

		if err := c.BodyParser(todo); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "cannot parse JSON",
			})
		}

		fmt.Println(todo)

		if (todo.Body == "") {
			return c.Status(400).JSON(fiber.Map{
				"error": "Missing todo body.",
			})
		}

		todos = append(todos, *todo)
		return c.Status(201).SendString(todo.BODY)
	})

	app.Get("/todos", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(todos)
	})

	app.Get("/todo/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		for _, todo := range todos {
			if todo.ID == id {
				return c.Status(200).JSON(todo)
			}
		}

		return c.Status(404).JSON(fiber.Map{
			"error": "todo not found",
		})
	})

	log.Fatal(app.Listen(":3000"))
}
