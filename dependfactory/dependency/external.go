package dependfactory

import (
	"github.com/ShaileshSurya/go-external-mocking/dependfactory/interfacer"
	dep "github.com/ShaileshSurya/go-external-mocking/external"
)

// External ...
type External struct {
	interfacer.External
}

// DependencyFunction ...
func (e External) DependencyFunction() error {
	return dep.DependencyFunction()
}
