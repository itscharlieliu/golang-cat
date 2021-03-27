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

func TestCatLineNumbers(t *testing.T) {

	commandOutputString := pkg.RunCmd("go", []string{"run", "../cmd/cat.go", "-n", "test.txt"})

	realOutputString := pkg.RunCmd("cat", []string{"-n", "test.txt"})

	if commandOutputString != realOutputString {
		t.Errorf("Expected: \n%s\nGot \n%s\n", realOutputString, commandOutputString)
	}
}

func TestCatCreateNewFile(t *testing.T) {
	pkg.RunCmd("go", []string{"run", "../cmd/cat.go", "test.txt", ">", "newfile.txt"})

	pkg.RunCmd("cat", []string{"test.txt", ">", "newfile2.txt"})

	text1 := pkg.RunCmd("cat", []string{"newfile.txt"})
	text2 := pkg.RunCmd("cat", []string{"newfile2.txt"})

	if text1 != text2 {
		t.Errorf("Expected: \n%s\nGot \n%s\n", text2, text1)
	}
}
