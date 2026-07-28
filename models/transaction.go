package models

import (
	"time"
)

type Transaction struct {
	ID          uint            `json:"id"`
	Ref         string          `json:"ref"`
	Date        string          `json:"date"`
	Description string          `json:"desc"`
	Category    string          `json:"category"`
	Amount      float64         `json:"amount"`
	Type        string          `json:"type"`
	Status      string          `json:"status"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
}

type Stats struct {
	TotalIncome  float64 `json:"total_income"`
	TotalExpense float64 `json:"total_expense"`
	Balance      float64 `json:"balance"`
	Savings      float64 `json:"savings"`
}

type CategoryStats struct {
	Category string  `json:"category"`
	Total    float64 `json:"total"`
	Count    int     `json:"count"`
}

type CreateTransactionInput struct {
	Date        string  `json:"date" binding:"required"`
	Description string  `json:"desc" binding:"required"`
	Category    string  `json:"category" binding:"required"`
	Amount      float64 `json:"amount" binding:"required"`
	Type        string  `json:"type" binding:"required"`
	Status      string  `json:"status"`
}

type UpdateTransactionInput struct {
	Date        string  `json:"date"`
	Description string  `json:"desc"`
	Category    string  `json:"category"`
	Amount      float64 `json:"amount"`
	Type        string  `json:"type"`
	Status      string  `json:"status"`
}