package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Todo struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Completed bool               `json:"completed"`
	Body      string             `json:"body"`
}

var collection *mongo.Collection

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	PORT := os.Getenv("PORT")
	MONGODB_URI := os.Getenv("MONGODB_URI")
	fmt.Println("using port: ", PORT)

	// Database connection

	clientOptions := options.Client().ApplyURI(MONGODB_URI)
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}

	defer client.Disconnect(context.Background())

	err = client.Ping(context.Background(), nil)

	if err != nil {
		log.Fatalf("Error pinging MongoDB: %v", err)
	}

	log.Println("Connected to MongoDB!")

	collection = client.Database("golang_db").Collection("todos")

	todos := []Todo{}

	app := fiber.New()

	// routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("Hello, World!")
	})

	app.Get("/todos", func(c *fiber.Ctx) error {
		var todos []Todo
		cursor, err := collection.Find(context.Background(), bson.M{})
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Error fetching todos from database",
			})
		}
		defer cursor.Close(context.Background())

		for cursor.Next(context.Background()) {
			var todo Todo
			if err := cursor.Decode(&todo); err != nil {
				return c.Status(500).JSON(fiber.Map{
					"error": "Error decoding todo from database",
				})
			}
			todos = append(todos, todo)
		}

		if err := cursor.Err(); err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Cursor error",
			})
		}

		return c.Status(200).JSON(todos)
	})

	app.Post("/todo", func(c *fiber.Ctx) error {
		todo := new(Todo)

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

	log.Fatal(app.Listen(":" + PORT))
}
