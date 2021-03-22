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

func ReadFile(filename string, showLineNumbers bool) {
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for i := 0; scanner.Scan(); i++ {
		if showLineNumbers {
			fmt.Printf("%d	", i+1)
		}
		fmt.Println(scanner.Text())
	}

	err = scanner.Err()
	check(err)

	// for {
	// 	// Create bytes slice of size 8
	// 	bytesSlice := make([]byte, SLICE_SIZE)
	// 	// Read file up to len(bytesSlice) bytes
	// 	numRead, err := file.Read(bytesSlice)
	// 	if err != nil {
	// 		// Either there was an error or we reached EOF
	// 		break
	// 	}
	// 	if (showLineNumbers) {
	// 		for i := 0; i < SLICE_SIZE; i++ {
	// 			if bytesSlice[i] == '\n' {

	// 			}
	// 		}
	// 	}
	// 	fmt.Printf("%s", string(bytesSlice[:numRead]))
	// }
}
