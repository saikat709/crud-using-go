package main

import (
	"context"
	"fmt"
	"log"
	"os"

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

		_, err := collection.InsertOne(context.Background(), todo)

		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Error inserting todo into database",
			})
		}

		return c.Status(201).SendString("Todo created successfully")
	})

	app.Get("/todo/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		objectID, err := primitive.ObjectIDFromHex(id)

		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Invalid ID format",
			})
		}

		filter := bson.M{"_id": objectID}
		var todo Todo
		err = collection.FindOne(context.Background(), filter).Decode(&todo)

		if err != nil {
			if err == mongo.ErrNoDocuments {
				return c.Status(404).JSON(fiber.Map{
					"error": "todo not found",
				})
			}
			return c.Status(500).JSON(fiber.Map{
				"error": "Error fetching todo from database",
			})
		}

		return c.Status(200).JSON(todo)
	})

	app.Delete("/todo/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		objectId, err := primitive.ObjectIDFromHex(id)

		if err != nil {
			log.Fatalf("Error converting string to int: %v", err)
			return c.Status(400).JSON(fiber.Map{
				"error": "todo not found.",
			})
		}

		filter := bson.M{"_id": objectId}

		_, err = collection.DeleteOne(context.Background(), filter)

		if err != nil {
			if err == mongo.ErrNoDocuments {
				return c.Status(404).JSON(fiber.Map{
					"error": "todo not found.",
				})
			}
			return c.Status(500).JSON(fiber.Map{
				"error": "Error deletibg todo from database.",
			})
		}

		return c.Status(200).JSON(fiber.Map{
			"message": "Deleted successfully.",
		})
	})

	app.Put("/todo/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		objectId, err := primitive.ObjectIDFromHex(id)
		todo := new((Todo))

		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Error converting hex to objectid.",
			})
		}

		if err := c.BodyParser(todo); err != nil {
			fmt.Println("BodyParser error:", err)
			return c.Status(400).JSON(fiber.Map{
				"error": "cannot parse JSON",
			})
		}

		filter := bson.M{"_id": objectId}
		update := bson.M{"$set": bson.M{
			"completed": todo.Completed,
			"body":      todo.Body,
		}}

		_, err = collection.UpdateOne(context.Background(), filter, update)

		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Error Updating.",
			})
		}

		return c.Status(404).JSON(fiber.Map{
			"message": "Updated succesfully.",
		})
	})

	log.Fatal(app.Listen(":" + PORT))
}
