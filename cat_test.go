package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestCat(t *testing.T) {

	stdout := os.Stdout // Keep backup of real stdout
	reader, writer, _ := os.Pipe()
	os.Stdout = writer // Set stdout to our writer that we just created

	readFile("test.txt")

	// Reset back to the normal state
	writer.Close()
	commandOutput, _ := ioutil.ReadAll(reader)
	os.Stdout = stdout

	commandOutputString := string(commandOutput)

	if commandOutputString != "test" {
		t.Errorf("Expected %s, got %s", "test", commandOutputString)
	}
}
