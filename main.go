package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"wallet-api/routes"
)

func main() {
	r := gin.Default() // Создаем сервер

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},                   // Указываем адрес фронтенда
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},            // Разрешенные методы
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // Разрешенные заголовки
		AllowCredentials: true,
	}))

	routes.SetupRoutes(r)

	r.Run(":8081")
}
