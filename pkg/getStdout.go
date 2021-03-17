package pkg

import (
	"io/ioutil"
	"os"
)

func GetStdout(callback func()) string {
	stdout := os.Stdout // Keep backup of real stdout
	reader, writer, _ := os.Pipe()
	os.Stdout = writer // Set stdout to our writer that we just created

	callback()

	// Reset back to the normal state
	writer.Close()
	commandOutput, _ := ioutil.ReadAll(reader)
	os.Stdout = stdout

	commandOutputString := string(commandOutput)

	return commandOutputString
}
