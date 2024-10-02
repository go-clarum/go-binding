package itests

import (
	clarum "github.com/go-clarum/go-binding"
	"os"
	"testing"
)

var testClient = clarum.Http().Client().
	Name("testClient").
	BaseUrl("http://localhost:8083/myApp").
	Timeout(2).
	Build()

var firstTestServer = clarum.Http().Server().
	Name("firstTestServer").
	Port(8083).
	Build()

var secondTestServer = clarum.Http().Server().
	Name("secondTestServer").
	Port(8084).
	Build()

func TestMain(m *testing.M) {
	result := m.Run()
	os.Exit(result)
}
