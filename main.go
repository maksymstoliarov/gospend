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

	// init storage
	// TODO: handle ~ somehow
	storage := expense.NewStorage("/home/max/.gospend/expenses.json")

	// init service
	service := expense.NewService(storage)

	// init params handler (read, parse, validate all parameters)
	// parse all params
	// switch available operations
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) == 0 {
		log.Fatal("no comand provided")
	}

	command := argsWithoutProg[0]
	switch command {
	case "list":
		service.List()
	case "add":
		// parse other arguments and validate
		validationErrors := make([]string, 0)

		// trick to parse args after comand
		originalArgs := os.Args
		os.Args = originalArgs[1:]

		flagAmount := flag.Float64("amount", 0, "expense amount")
		flagCategory := flag.String("category", "", "expense category")
		flagNote := flag.String("note", "", "expense note")
		flag.Parse()

		// rollback args to original value
		os.Args = originalArgs

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
	default:
		log.Fatalf("unknown command: %s", command)
	}

	// init service component, based on parameter and pass all parameters into it

	// save result into json file
}
