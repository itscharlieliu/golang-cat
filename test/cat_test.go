package test

import (
	"io/ioutil"
	"os"
	"os/exec"
	"testing"

	"github.com/itscharlieliu/golang-cat/pkg"
)

func TestCat(t *testing.T) {

	commandOutputString := pkg.GetStdout(func() { pkg.ReadFile("test.txt") })

	exe, _ := exec.LookPath("cat")

	reader, writer, _ := os.Pipe()

	cmd := &exec.Cmd{
		Path:   exe,
		Args:   []string{exe, "test.txt"},
		Stdout: writer,
		Stderr: writer,
	}

	cmd.Run()

	writer.Close()
	output, _ := ioutil.ReadAll(reader)

	realOutputString := string(output)

	if commandOutputString != realOutputString {
		t.Errorf("Expected %s, got %s", realOutputString, commandOutputString)
	}
}
