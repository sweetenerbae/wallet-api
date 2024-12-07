package main

import (
	"github.com/gin-gonic/gin"
	"wallet-api/routes"
)

func main() {
	r := gin.Default() // Создаем сервер

	routes.SetupRoutes(r)

	r.Run(":8081")
}
