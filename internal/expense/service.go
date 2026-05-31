package expense

import (
	"fmt"
	"log"
	"time"
)

var monthFilterFormat = "2006-01"

type Service struct {
	storage *Storage
}

func NewService(s *Storage) *Service {
	return &Service{
		storage: s,
	}
}

func (s *Service) List(amount float64, category string, note string, month string) {
	var expenses []Expense
	if amount != 0 || category != "" || note != "" || month != "" {
		var monthTime *time.Time
		if month != "" {
			var err error
			monthTime, err = s.parseMonthFilter(month)
			if err != nil {
				log.Fatal("error getting list of expenses. " + err.Error())
			}
		}

		expenses = make([]Expense, 0)
		for _, e := range s.storage.GetAll() {
			ok := true

			ok = ok && (amount != 0 && amount == e.Amount || amount == 0)
			ok = ok && (category != "" && category == e.Category || category == "")
			ok = ok && (note != "" && note == e.Note || note == "")
			ok = ok && (month != "" && monthTime != nil && monthTime.Format(monthFilterFormat) == e.CreatedAt.Format(monthFilterFormat) || month == "")

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

func (s *Service) Total(month string) {
	var expenses []Expense
	if month != "" {
		var monthTime *time.Time
		if month != "" {
			var err error
			monthTime, err = s.parseMonthFilter(month)
			if err != nil {
				log.Fatal("error getting total of expenses. " + err.Error())
			}
		}

		expenses = make([]Expense, 0)
		for _, e := range s.storage.GetAll() {
			if monthTime != nil && e.CreatedAt.Format(monthFilterFormat) == monthTime.Format(monthFilterFormat) {
				expenses = append(expenses, e)
			}
		}
	} else {
		expenses = s.storage.GetAll()
	}

	fmt.Println("Total:", len(expenses))
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

func (s *Service) parseMonthFilter(month string) (*time.Time, error) {
	if month == "" {
		return nil, fmt.Errorf("month is empty")
	}

	monthTime, err := time.Parse(monthFilterFormat, month)
	if err != nil {
		return nil, fmt.Errorf("wrong format of month filter")
	}

	return &monthTime, err
}
