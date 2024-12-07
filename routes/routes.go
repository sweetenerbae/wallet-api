package routes

import (
	"github.com/gin-gonic/gin"
	"wallet-api/handlers"
)

func SetupRoutes(r *gin.Engine) {
	//Пользователи
	r.POST("/users", handlers.CreateUser)
	r.GET("/users/:id", handlers.GetUserBalance)

	r.GET("/users", handlers.GetAllUsers)

	//Транзакции
	r.POST("/transactions", handlers.CreateTransaction)
	r.GET("/transactions/:user_id", handlers.GetUserTransactions)
}
