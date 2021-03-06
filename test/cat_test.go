package test

import (
	"testing"

	"github.com/itscharlieliu/golang-cat/pkg"
)

func TestCatNoArgs(t *testing.T) {

	commandOutputString := pkg.RunCmd("go", []string{"run", "../cmd/cat.go", "test.txt"})

	realOutputString := pkg.RunCmd("cat", []string{"test.txt"})

	if commandOutputString != realOutputString {
		t.Errorf("Expected %s, got %s", realOutputString, commandOutputString)
	}
}

func TestCatMultipleFiles(t *testing.T) {

	commandOutputString := pkg.RunCmd("go", []string{"run", "../cmd/cat.go", "test.txt", "test2.txt"})

	realOutputString := pkg.RunCmd("cat", []string{"test.txt", "test2.txt"})

	if commandOutputString != realOutputString {
		t.Errorf("Expected %s, got %s", realOutputString, commandOutputString)
	}
}

func TestCatSuppressEmptyLines(t *testing.T) {

	commandOutputString := pkg.RunCmd("go", []string{"run", "../cmd/cat.go", "-s", "test.txt"})

	realOutputString := pkg.RunCmd("cat", []string{"-s", "test.txt"})

	if commandOutputString != realOutputString {
		t.Errorf("Expected: \n%s\nGot \n%s\n", realOutputString, commandOutputString)
	}
}

func TestCatHighlightEndOfLine(t *testing.T) {

	commandOutputString := pkg.RunCmd("go", []string{"run", "../cmd/cat.go", "-E", "test3.txt"})

	realOutputString := pkg.RunCmd("cat", []string{"-E", "test3.txt"})

	if commandOutputString != realOutputString {
		t.Errorf("Expected: \n%s\nGot \n%s\n", realOutputString, commandOutputString)
	}
}
