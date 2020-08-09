package main

import (
	"flag"
	"fmt"

	factory "github.com/ShaileshSurya/go-external-mocking/dependfactory"
)

func main() {
	env := flag.String("env", "dev", "Environment Test Or Development")
	flag.Parse()
	external := factory.GetExternal(*env)
	err := external.DependencyFunction()
	if err != nil {
		fmt.Println("Error from dependency function")
	}
}
