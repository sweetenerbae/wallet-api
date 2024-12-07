package models

// Пользователь с ID и балансом
type User struct {
	ID      int     `json:"id"`
	Balance float64 `json:"balance"`
}

// Транзакция (пополнение или списание)
type Transaction struct {
	ID        int     `json:"id"`
	UserID    int     `json:"user_id"`
	Amount    float64 `json:"amount"` // Положительное для пополнения, отрицательное для списания
	Timestamp string  `json:"timestamp"`
}
