## Necessary tools:
1. To load env file
    ```bash
    go get github.com/joho/godotenv
    ```

## Running the database
1. Write the compose.yml for mongodb container
2. install docker if not installed -
    ```bash
    sudo snap install docker
    ```
3. run the container
    ```bash
    sudo docker compose up -d
    ```
4. check running containers
    ```bash
    sudo docker ps
    ```

The url of mongodb would be: `mongodb://admin:secret@localhost:27017/?authSource=admin`


## Connecting from GO
1. install mongo dependency - using json and for connecting monogdb
```
go get go.mongodb.org/mongo-driver/bson
go get go.mongodb.org/mongo-driver/bson/primitive
go get go.mongodb.org/mongo-driver/mongo
go get go.mongodb.org/mongo-driver/mongo/options
```
