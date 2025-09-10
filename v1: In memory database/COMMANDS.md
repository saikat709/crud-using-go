# Related Commands

1. Creating a project
    ```bash
    go mod init crud-go
    ```

2. installing a package to the project
    ```bash
    go get <package name>

    // go get github.com/gofiber/fiber/v2
    ```

3. running the projects main file
    ```bash
    go run main.go
    ```

4. For auto restart we use air.
    ```bash

    go install github.com/air-verse/air@latest

    echo 'export PATH=$PATH:$HOME/go/bin' >> ~/.bashrc

    source ~/.bashrc

    // write a .air.toml file
    air

    ```