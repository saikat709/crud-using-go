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

- **URL:** `/`
- **Method:** `GET`
- **Description:** Returns a "Hello, World!" message.
- **Success Response:**
  - **Code:** 200 OK
  - **Content:** `Hello, World!`
