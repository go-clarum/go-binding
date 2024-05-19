package itests

import (
	go_binding "github.com/go-clarum/go-binding"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	defer go_binding.Shutdown()

	result := m.Run()

	os.Exit(result)
}
