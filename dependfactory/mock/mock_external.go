package mock

import (
	"errors"

	"github.com/ShaileshSurya/go-external-mocking/dependfactory/interfacer"
)

// MockExternal ...
type MockExternal struct {
	interfacer.External
}

// DependencyFunction ...
func (e MockExternal) DependencyFunction() error {
	return errors.New("Error Returning from depenedency function")
}
