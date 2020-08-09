package dependfactory

import (
	dep "github.com/ShaileshSurya/go-external-mocking/dependfactory/dependency"
	"github.com/ShaileshSurya/go-external-mocking/dependfactory/interfacer"
	"github.com/ShaileshSurya/go-external-mocking/dependfactory/mock"
)

// GetExternal ...
func GetExternal(env string) interfacer.External {
	if env == "test" {
		return mock.MockExternal{}
	}
	return dep.External{}
}
