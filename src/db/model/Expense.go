package model

import (
	"time"
)

type ExpenseType string

const (
	ExpenseTypeIncome  ExpenseType = "income"
	ExpenseTypeExpense ExpenseType = "expense"
)

type Expense struct {
	ID       int         `json:"ID,omitempty"`
	Date     time.Time   `json:"date"`
	Amount   float64     `json:"amount,omitempty"`
	Type     ExpenseType `json:"type,omitempty"`
	Category string      `json:"category,omitempty"`
}
