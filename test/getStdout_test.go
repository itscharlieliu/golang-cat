package test

import (
	"fmt"
	"testing"

	"github.com/itscharlieliu/golang-cat/pkg"
)

func testFunc(str string) {
	fmt.Print(str)
}

func TestGetStdout(t *testing.T) {
	result := pkg.GetStdout(func() { testFunc("hello world") })

	if result != "hello world" {
		t.Errorf("Expected %s, got %s", "hello world", result)
	}
}
