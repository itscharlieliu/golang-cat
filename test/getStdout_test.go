package main

import (
	"fmt"
	"testing"
)

func testFunc(str string) {
	fmt.Print(str)
}

func TestGetStdout(t *testing.T) {
	result := getStdout(func() { testFunc("hello world") })

	if result != "hello world" {
		t.Errorf("Expected %s, got %s", "test", result)
	}
}
