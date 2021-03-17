package pkg

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadFile(filename string) {
	file, err := os.Open(filename)
	check(err)

	for {
		// Create bytes slice of size 8
		bytesSlice := make([]byte, 8)
		// Read file up to len(bytesSlice) bytes
		numRead, err := file.Read(bytesSlice)
		if err != nil {
			// Either there was an error or we reached EOF
			break
		}
		fmt.Printf("%s", string(bytesSlice[:numRead]))
	}
}
