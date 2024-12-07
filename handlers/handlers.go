package handlers

import (
	"net/http"
	"strconv"
	"time"
	"wallet-api/models"

	"github.com/gin-gonic/gin"
)

// Хранилище данных
var users = make(map[int]*models.User)
var transactions = []models.Transaction{}
var userIDCounter = 1
var transactionIDCounter = 1

// Создание нового пользователя
func CreateUser(c *gin.Context) {
	var input struct {
		Name    string  `json:"name"`
		Balance float64 `json:"balance"`
	}

	//Привязка данных из тела запроса к переменной input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := models.User{
		ID:      userIDCounter,
		Balance: input.Balance,
	}
	users[userIDCounter] = &user
	userIDCounter++

	c.JSON(http.StatusCreated, user)
}

// Получение баланса пользователя
func GetUserBalance(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, exists := users[id]

	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id, "balance": user.Balance})
}

// Создание транзакции
func CreateTransaction(c *gin.Context) {
	var input struct {
		UserID int     `json:"user_id"`
		Amount float64 `json:"amount"`
	}

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, exists := users[input.UserID]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	//проверяем баланс при списании
	if input.Amount < 0 && user.Balance+input.Amount < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Insufficient balance"})
		return
	}

	//Обновляем баланс и добавляем транзакцию
	user.Balance += input.Amount
	transaction := models.Transaction{
		ID:        transactionIDCounter,
		UserID:    input.UserID,
		Amount:    input.Amount,
		Timestamp: time.Now().Format(time.RFC3339),
	}
	transactions = append(transactions, transaction)
	transactionIDCounter++

	c.JSON(http.StatusOK, transaction)
}

// Получение транзакций пользователя
func GetUserTransactions(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Param("user_id"))
	var userTransactions []models.Transaction

	for _, t := range transactions {
		if t.UserID == userID {
			userTransactions = append(userTransactions, t)
		}
	}

	if len(userTransactions) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No transactions found"})
		return
	}

	c.JSON(http.StatusOK, userTransactions)
}

// Получение всех пользователей
func GetAllUsers(c *gin.Context) {
	// Преобразуем карту пользователей в срез
	userList := []models.User{}
	for _, user := range users {
		userList = append(userList, *user)
	}

	// Возвращаем срез пользователей
	if len(userList) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No users found"})
		return
	}

	c.JSON(http.StatusOK, userList)
}
