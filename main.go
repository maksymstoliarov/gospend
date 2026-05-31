package main

import (
	"fmt"
	"time"

	"github.com/maksymstoliarov/gospend/internal/expense"
)

func main() {
	fmt.Println("Hello gospend")

	// init storage
	_ = expense.NewStorage("~/.gospend/expenses.json")

	// init params handler (read, parse, validate all parameters)

	// init service component, based on parameter and pass all parameters into it
	newExpense := expense.Expense{
		ID:        1,
		Amount:    9.99,
		Category:  "food",
		Note:      "groceries",
		CreatedAt: time.Date(2026, 1, 13, 12, 0, 0, 0, time.Local),
	}

	fmt.Println(newExpense)

	// save result into json file
}
