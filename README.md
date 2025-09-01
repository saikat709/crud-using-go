# CRUD Go

This is a simple CRUD application built with Go and the [Fiber](https://gofiber.io/) web framework.

## Prerequisites

- [Go](https://golang.org/doc/install) (version 1.22.2 or higher)

## Getting Started

1. **Clone the repository:**

   ```bash
   git clone https://github.com/saikat709/crud-go.git
   cd crud-go
   ```

2. **Install dependencies:**

   ```bash
   go mod tidy
   ```

3. **Run the application:**

   ```bash
   go run main.go
   ```
   
   The application will be running at `http://localhost:3000`.


4. Extra:  **To run in watch mode**

   ```bash
   go install github.com/air-verse/air@latest
   echo 'export PATH=$PATH:$HOME/go/bin' >> ~/.bashrc
   source ~/.bashrc
   // write a .air.toml file
   air

   ```

## API Endpoints

### Root


### Create a Todo
- **URL:** `/api/todo`
- **Method:** `POST`
- **Content-Type:** `application/json`
- **Body Example:**
   ```json
   {
      "id": 1,
      "completed": false,
      "body": "New Body"
   }
   ```
- **Curl Example:**
   ```bash
   curl -X POST http://localhost:3000/api/todo \
      -H "Content-Type: application/json" \
      -d '{"id":1,"completed":false,"body":"New Body"}'
   ```

### Get All Todos
- **URL:** `/api/todos`
- **Method:** `GET`
- **Curl Example:**
   ```bash
   curl http://localhost:3000/api/todos
   ```

### Get Todo by ID
- **URL:** `/todo/:id`
- **Method:** `GET`
- **Curl Example:**
   ```bash
   curl http://localhost:3000/todo/1
   ```
