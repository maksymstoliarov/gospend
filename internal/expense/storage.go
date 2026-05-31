package expense

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
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
	if s.content, err = os.ReadFile(p); err == nil {
		// file exists, check for corruption

		// validate
		if valid := json.Valid(s.content); !valid {
			log.Fatal("error reading storage file. file does not contain valid JSON")
		}
	} else if os.IsNotExist(err) {
		// file does not exist, create empty file
		// ensure dir
		dirName := filepath.Dir(p)
		if _, err = os.Stat(dirName); err != nil {
			// dir does not exist
			if err = os.MkdirAll(dirName, os.ModePerm); err != nil {
				log.Fatal("error creating storage file dir. " + err.Error())
			}
		}

		// write empty array to the file
		if err = os.WriteFile(p, []byte("[]"), os.ModePerm); err != nil {
			log.Fatal("error writing initial value to file. " + err.Error())
		}
	} else if os.IsPermission(err) {
		// file has corrupted permission, log.Fatal
		log.Fatal("error accessing storage file. permission denied")
	} else {
		log.Fatal("error reading storage file. " + err.Error())
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

func (s *Storage) GetOneWithMaxID() Expense {
	// unmarshal json content, loop and find highest id
	var maxID, maxI int
	expenses := s.unmarshalExpenses()
	for i, e := range expenses {
		if e.ID > maxID {
			maxID = e.ID
			maxI = i
		}
	}
	if maxID == 0 {
		return Expense{}
	}
	return expenses[maxI]
}

func (s *Storage) Add(newExpense Expense) {
	// unmarshal json content, append one json object
	expenses := s.unmarshalExpenses()
	expenses = append(expenses, newExpense)

	s.marshalExpenses(expenses)

	if err := os.WriteFile(s.path, s.content, os.ModePerm); err != nil {
		log.Fatal("error writing added expense into file. " + err.Error())
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

	s.marshalExpenses(expenses)

	if err := os.WriteFile(s.path, s.content, os.ModePerm); err != nil {
		log.Fatal("error writing deleted expense into file. " + err.Error())
	}

	return nil
}

func (s *Storage) DeleteAll() {
	s.content = []byte("[]")
	if err := os.WriteFile(s.path, s.content, os.ModePerm); err != nil {
		log.Fatal("error writing deleted expenses into file. " + err.Error())
	}
}

func (s *Storage) unmarshalExpenses() []Expense {
	result := make([]Expense, 0)
	if err := json.Unmarshal(s.content, &result); err != nil {
		log.Fatal("error getting all expenses. " + err.Error())
	}
	return result
}

func (s *Storage) marshalExpenses(expenses []Expense) {
	var err error
	s.content, err = json.Marshal(expenses) //json.MarshalIndent(expenses, "", "    ")
	if err != nil {
		log.Fatal("error saving added expense. " + err.Error())
	}
}
