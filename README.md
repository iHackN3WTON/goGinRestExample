# SPIKE: Go Gin REST Example

## Instructions

- Install docker desktop, download the installer from this URL: https://hub.docker.com/
- Install GoLang, download the installer from this URL: https://go.dev/dl/
- Install VSCode, dowload the installer from this URL: https://code.visualstudio.com/download
- Install Postman, download the installer from this URL: https://www.postman.com/downloads/
- Clone the repository and open the cloned folder with VSCode

- Create the docker image running the following command in a terminal:

        docker run --name GoPostgreSQL -e POSTGRES_PASSWORD=docker -e POSTGRES_DB=go_db -p 5432:5432 -d postgres

- Download the Gin repository and their dependencies with the following commands in a terminal:

        go get -u github.com/gin-gonic/gin
        go get -u github.com/go-pg/pg
        go get -u github.com/google/uuid
        go get -u github.com/gin-contrib/cors

- Run the code with the following command in a terminal:

        go run main.go

- Interact with the API with these end point using postman:
    - GET /todos
    - POST /todo
        body:
        {
            "title": "any title",
            "body": "any body",
            "completed": "false"
        }
    - GET /todo/:id
    - PUT /todo/:id
        body:
        {
            "completed": "true"
        }
    - DELETE /todo/:id
