package main

import (
	"log"

	config "github.com/iHackN3WTON/goGinRestExample/configs"
	routes "github.com/iHackN3WTON/goGinRestExample/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.Connect()

	//init router
	router := gin.Default()

	//CORS
	router.Use(cors.Default())

	//Router Handlers / Endpoints
	routes.Routes(router)

	log.Fatal(router.Run(":4747"))
}
