package mock

import (
	"errors"
	"fmt"

	"github.com/ShaileshSurya/go-external-mocking/dependfactory/interfacer"
)

// MockExternal ...
type MockExternal struct {
	interfacer.External
}

// DependencyFunction ...
func (e MockExternal) DependencyFunction() error {
	fmt.Println("Hurray Mock Dependency Fuction is called...")
	return errors.New("Error Returning from depenedency function")
}
