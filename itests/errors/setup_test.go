package errors

import (
	"errors"
	clarum "github.com/go-clarum/go-binding"
	"os"
	"strings"
	"testing"
)

var errorsClient = clarum.Http().Client().
	Name("errorsClient").
	Timeout(2).
	Build()

var errorsServer = clarum.Http().Server().
	Name("errorsServer").
	Port(8083).
	Build()

func TestMain(m *testing.M) {
	result := m.Run()
	os.Exit(result)
}

func checkErrors(t *testing.T, expectedErrors []string, actionErrors ...error) {
	allErrors := errors.Join(actionErrors...)

	if allErrors == nil {
		t.Error("One error expected, but there was none.")
	} else {
		for _, value := range expectedErrors {
			if !strings.Contains(allErrors.Error(), value) {
				t.Errorf("Unexpected errors: %s", allErrors)
			}
		}
	}
}
