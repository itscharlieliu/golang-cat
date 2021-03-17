package test

import (
	"testing"

	"github.com/itscharlieliu/golang-cat/pkg"
)

func TestCat(t *testing.T) {

	commandOutputString := pkg.GetStdout(func() { main.ReadFile("test.txt") })

	if commandOutputString != "test" {
		t.Errorf("Expected %s, got %s", "test", commandOutputString)
	}
}
