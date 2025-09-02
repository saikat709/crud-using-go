package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	ID        int    `json:"id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

func main() {
	fmt.Println("hello world")

	app := fiber.New()

	todos := []Todo{}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("Hello, World!")
	})

	app.Get("/todos", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(todos)
	})

	app.Post("/todo", func(c *fiber.Ctx) error {
		todo := &Todo{}

		if err := c.BodyParser(todo); err != nil {
			fmt.Println("BodyParser error:", err)
			return c.Status(400).JSON(fiber.Map{
				"error": "cannot parse JSON",
			})
		}

		fmt.Printf("Received Todo: %+v\n", todo)

		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{
				"error": "Missing todo body.",
			})
		}

		todo.ID = len(todos) + 1
		todos = append(todos, *todo)

		return c.Status(201).SendString("Todo created successfully")
	})

	app.Get("/todo/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))

		if err != nil {
			log.Fatalf("Error converting string to int: %v", err)
		}

		for _, todo := range todos {
			if todo.ID == id {
				return c.Status(200).JSON(todo)
			}
		}

		return c.Status(404).JSON(fiber.Map{
			"error": "todo not found",
		})
	})

	app.Delete("/todo/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))

		if err != nil {
			log.Fatalf("Error converting string to int: %v", err)
		}

		for index, todo := range todos {
			if todo.ID == id {
				todos = append(todos[:index], todos[index+1:]...)
				return c.Status(200).JSON(fiber.Map{
					"message": "todo deleted successfully",
				})
			}
		}

		return c.Status(404).JSON(fiber.Map{
			"error": "todo not found",
		})
	})

	app.Put("/todo/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))

		if err != nil {
			log.Fatalf("Error converting string to int: %v", err)
		}

		updatedTodo := &Todo{}

		if err := c.BodyParser(updatedTodo); err != nil {
			fmt.Println("BodyParser error:", err)
			return c.Status(400).JSON(fiber.Map{
				"error": "cannot parse JSON",
			})
		}

		for index, todo := range todos {
			if todo.ID == id {
				todos[index].Body = updatedTodo.Body
				todos[index].Completed = updatedTodo.Completed
				return c.Status(200).JSON(todos[index])
			}
		}

		return c.Status(404).JSON(fiber.Map{
			"error": "todo not found",
		})
	})

	app.Patch("/todo/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))

		if err != nil {
			log.Fatalf("Error converting string to int: %v", err)
		}

		updatedFields := &Todo{}

		if err := c.BodyParser(updatedFields); err != nil {
			fmt.Println("BodyParser error:", err)
			return c.Status(400).JSON(fiber.Map{
				"error": "cannot parse JSON",
			})
		}

		for index, todo := range todos {
			if todo.ID == id {
				if updatedFields.Body != "" {
					todos[index].Body = updatedFields.Body
				}
				todos[index].Completed = updatedFields.Completed
				return c.Status(200).JSON(todos[index])
			}
		}

		return c.Status(404).JSON(fiber.Map{
			"error": "todo not found",
		})
	})

	log.Fatal(app.Listen(":3000"))
}
