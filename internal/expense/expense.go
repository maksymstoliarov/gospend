package expense

import "time"

type Expense struct {
	ID        int       `json:"id"`
	Amount    float64   `json:"amount"`
	Category  string    `json:"category"`
	Note      string    `json:"note"`
	CreatedAt time.Time `json:"created_at"`
}
