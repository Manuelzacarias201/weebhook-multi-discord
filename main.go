package main

import (
	"weebhook/infraestructure"
	"weebhook/infraestructure/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	webhookHandler, reviewHandler := infraestructure.Init()

	router := gin.Default()

	routes.Routes(router, webhookHandler, reviewHandler)

	router.Run(":8080")
}
