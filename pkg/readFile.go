package pkg

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

const SLICE_SIZE = 8

func writeFlag(write bool) string {
	if write {
		return "$"
	}
	return ""
}

func ReadFile(filename string, showLineNumbers bool, hideEmptyLines bool, highlightEndOfString bool) {
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	reader := bufio.NewReader(file)

	continueReading := true

	for i, emptyLines := 0, 0; continueReading; i++ {
		line, err := reader.ReadString('\n')

		if err != nil {
			// Either there was an error or we reached EOF
			continueReading = false
		}

		if line == "\n" {
			emptyLines++
		} else {
			emptyLines = 0
		}

		if emptyLines > 1 {
			continue
		}

		if showLineNumbers {
			fmt.Printf("     %d	", i+1)
		}

		lineSize := len(line)
		if lineSize > 0 && line[lineSize-1] == '\n' {
			if lineSize == 1 {
				print(writeFlag(highlightEndOfString) + "\n")
				continue
			}
			fmt.Print(line[:lineSize-2] + writeFlag(highlightEndOfString) + string(line[lineSize-1]))
		}
	}
}
