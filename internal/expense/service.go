package expense

import (
	"fmt"
	"time"
)

type Service struct {
	storage *Storage
}

func NewService(s *Storage) *Service {
	return &Service{
		storage: s,
	}
}

func (s *Service) List() {
	expenses := s.storage.GetAll()
	fmt.Println(expenses)
}

func (s *Service) Add(amount float64, category string, note string) {
	expenseWithMaxID := s.storage.GetOneWithMaxID()
	var id int
	if expenseWithMaxID.ID != 0 {
		id = expenseWithMaxID.ID + 1
	} else {
		id = 1
	}
	newExpense := Expense{id, amount, category, note, time.Now()}
	s.storage.Add(newExpense)
	fmt.Println("new expense is added")
}

func (s *Service) Clear() {
	s.storage.DeleteAll()
	fmt.Println("expenses are cleared")
}
