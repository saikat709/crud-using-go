package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	Id        int    `json:"id"`
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

	app.Post("/api/todo", func(c *fiber.Ctx) error {
		todo := &Todo{}

		if err := c.BodyParser(todo); err != nil {
			fmt.Println("BodyParser error:", err)
			return c.Status(400).JSON(fiber.Map{
				"error": "cannot parse JSON",
			})
		}

		fmt.Printf("Received Todo: %+v\n", todo)

		// if (todo.Body == "") {
		// 	return c.Status(400).JSON(fiber.Map{
		// 		"error": "Missing todo body.",
		// 	})
		// }

		todo.Id = len(todos) + 1
		todos = append(todos, *todo)

		fmt.Printf("fuck")
		return c.Status(201).SendString("Todo created successfully")
	})

	app.Get("/api/todos", func(c *fiber.Ctx) error {
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
