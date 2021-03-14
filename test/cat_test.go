package main

import (
	"testing"
)

func TestCat(t *testing.T) {

	commandOutputString := getStdout(func() { readFile("test.txt") })

	if commandOutputString != "test" {
		t.Errorf("Expected %s, got %s", "test", commandOutputString)
	}
}
