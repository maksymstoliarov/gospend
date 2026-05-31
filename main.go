package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/maksymstoliarov/gospend/internal/expense"
)

func main() {
	fmt.Println("Hello gospend")

	// trick to parse args after comand
	originalArgs := os.Args
	os.Args = originalArgs[1:]

	if len(os.Args) == 0 {
		log.Fatal("no arguments provided")
	}

	flagAmount := flag.Float64("amount", 0, "expense amount")
	flagCategory := flag.String("category", "", "expense category")
	flagNote := flag.String("note", "", "expense note")
	flagMonth := flag.String("month", "", "filter by Y-m (example: 2026-05)")
	flagID := flag.Int("id", 0, "expense id")
	flag.Parse()

	// rollback args to original value
	os.Args = originalArgs

	// init storage
	// TODO: handle ~ somehow
	storage := expense.NewStorage("/home/max/.gospend/expenses.json")

	// init service
	service := expense.NewService(storage)

	// init params handler (read, parse, validate all parameters)
	// parse all params
	// switch available operations

	command := os.Args[1]
	switch command {
	case "list":
		service.List(*flagAmount, *flagCategory, *flagNote, *flagMonth)
	case "add":
		// parse other arguments and validate
		validationErrors := make([]string, 0)

		if *flagAmount == 0.00 {
			validationErrors = append(validationErrors, "amount cannot be 0")
		}
		if *flagCategory == "" {
			validationErrors = append(validationErrors, "category cannot be empty")
		}
		if *flagNote == "" {
			validationErrors = append(validationErrors, "note cannot be empty")
		}

		if len(validationErrors) > 0 {
			log.Fatal("cannot add new expense: " + strings.Join(validationErrors, ", "))
		} else {
			service.Add(*flagAmount, *flagCategory, *flagNote)
		}
	case "clear":
		// TODO: add confirm
		service.Clear()
	case "total":
		service.Total(*flagMonth)
	case "delete":
		if *flagID == 0 {
			log.Fatal("cannot delete expense. id cannot be empty")
		}
		service.Delete(*flagID)
	default:
		log.Fatalf("unknown command: %s", command)
	}
}
