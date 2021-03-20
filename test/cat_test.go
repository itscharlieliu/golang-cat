package test

import (
	"testing"

	"github.com/itscharlieliu/golang-cat/pkg"
)

func TestCatNoArgs(t *testing.T) {

	commandOutputString := pkg.GetStdout(func() { pkg.ReadFile("test.txt") })

	realOutputString := pkg.RunCmd("cat", []string{"test.txt"})

	if commandOutputString != realOutputString {
		t.Errorf("Expected %s, got %s", realOutputString, commandOutputString)
	}
}

// func TestCatMultipleFiles(t *testing.T) {

// 	commandOutputString := pkg.GetStdout(func() { pkg.ReadFile("test.txt") })

// 	exe, _ := exec.LookPath("cat")

// 	reader, writer, _ := os.Pipe()

// 	cmd := &exec.Cmd{
// 		Path:   exe,
// 		Args:   []string{exe, "test.txt"},
// 		Stdout: writer,
// 		Stderr: writer,
// 	}

// 	cmd.Run()

// 	writer.Close()
// 	output, _ := ioutil.ReadAll(reader)

// 	realOutputString := string(output)

// 	if commandOutputString != realOutputString {
// 		t.Errorf("Expected %s, got %s", realOutputString, commandOutputString)
// 	}
// }
