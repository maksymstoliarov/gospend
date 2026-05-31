package expense

import (
	"encoding/json"
	"fmt"
	"os"
)

type Storage struct {
	path    string
	file    *os.File
	content []byte
}

func NewStorage(p string) *Storage {
	s := &Storage{
		path: p,
	}
	var err error
	// check if file exists
	if s.file, err = os.Open(p); err == nil {
		// file exists, check for corruption

		// read
		_, err := s.file.Read(s.content)
		if err != nil {
			panic("error reading storage file. " + err.Error())
		}

		// validate
		if valid := json.Valid(s.content); !valid {
			panic("error reading storage file. file does not contain valid JSON")
		}
	} else if os.IsNotExist(err) {
		// file does not exist, create empty file
		s.file, err = os.Create(p)
		if err != nil {
			panic("error creating storage file. " + err.Error())
		}
	} else if os.IsPermission(err) {
		// file has corrupted permission, panic
		panic("error accessing storage file. permission denied")
	} else {
		panic("error opening storage file. " + err.Error())
	}

	return s
}

func (s *Storage) GetAll() []Expense {
	// unmarshal json content and return slice of expenses
	return s.unmarshalExpenses()
}

func (s *Storage) GetOne(id int) (Expense, error) {
	// unmarshal json content and return only one expense
	for _, e := range s.unmarshalExpenses() {
		if e.ID == id {
			return e, nil
		}
	}
	return Expense{}, fmt.Errorf("expense not found")

}

func (s *Storage) Add(newExpense Expense) {
	// unmarshal json content, append one json object
	expenses := s.unmarshalExpenses()
	expenses = append(expenses, newExpense)

	var err error
	s.content, err = json.Marshal(expenses)
	if err != nil {
		panic("error saving added expense. " + err.Error())
	}
	if _, err = s.file.Write(s.content); err != nil {
		panic("error writing added expense into file. " + err.Error())
	}

}

func (s *Storage) Delete(id int) error {
	// unmarshal json content, delete json object
	var expenses []Expense
	for _, e := range s.unmarshalExpenses() {
		if e.ID != id {
			expenses = append(expenses, e)
		}
	}

	var err error
	s.content, err = json.Marshal(expenses)
	if err != nil {
		panic("error saving deleted expense. " + err.Error())
	}
	if _, err = s.file.Write(s.content); err != nil {
		panic("error writing deleted expense into file. " + err.Error())
	}

	return nil
}

func (s *Storage) unmarshalExpenses() []Expense {
	var result []Expense
	if err := json.Unmarshal(s.content, result); err != nil {
		panic("error getting all expenses. " + err.Error())
	}
	return result
}
