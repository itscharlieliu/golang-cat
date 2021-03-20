package pkg

import (
	"io/ioutil"
	"os"
	"os/exec"
)

func RunCmd(command string, args []string) string {

	exe, _ := exec.LookPath(command)

	reader, writer, _ := os.Pipe()

	cmdSlice := []string{exe}
	cmdArgs := append(cmdSlice, args...)

	cmd := &exec.Cmd{
		Path:   exe,
		Args:   cmdArgs,
		Stdout: writer,
		Stderr: writer,
	}

	cmd.Run()

	writer.Close()
	output, _ := ioutil.ReadAll(reader)

	return string(output)

}
