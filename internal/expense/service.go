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

func (s *Service) List(amount float64, category string, note string) {
	var expenses []Expense
	if amount != 0 || category != "" || note != "" {
		expenses = make([]Expense, 0)
		for _, e := range s.storage.GetAll() {
			ok := true

			ok = ok && (amount != 0 && amount == e.Amount || amount == 0)
			ok = ok && (category != "" && category == e.Category || category == "")
			ok = ok && (note != "" && note == e.Note || note == "")

			if ok {
				expenses = append(expenses, e)
			}
		}
	} else {
		expenses = s.storage.GetAll()
	}

	if len(expenses) == 0 {
		fmt.Println("no expenses found")
	} else {
		fmt.Println(expenses)
	}
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
