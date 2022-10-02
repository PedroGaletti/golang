package fileio

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

/*
	Exactly one of O_RDONLY, O-WRONLY, or ORDWR must be specified

	O_RDONLY: open the file read-only
	O_WRONLY: open the file write-only
	O_RDWR: open the file read-write

	these can be or'ed

	O_APPEND: append data to the file when writing
	O_CREATE: create a new file if none exists
	O_EXCL: used with O_CREATE, file must not exist
	O_SYNC: open for synchronous I/O
	O_TRUNC: truncate regular weitable file when opened
*/

func IfFileExistsWrite() {
	_, err := os.Stat("data.txt") // Returning the info of file
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("File doesn't exist")
	} else {
		f, err := os.OpenFile("data.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) // Opening the file with some flags and permission
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		if _, err := f.WriteString("13\n"); err != nil { // Writing new line in the file
			log.Fatal(err)
		}
	}
}

func Files() {
	f, err := os.Create("data.txt") // Creating an file
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close() // Close the file after we open it in the end of the func "defer"

	// Insert int values into string array
	iPrime := []int{2, 3, 5, 7, 11}
	var sPrime []string
	for _, i := range iPrime {
		sPrime = append(sPrime, strconv.Itoa(i))
	}

	for _, num := range sPrime {
		_, err := f.WriteString(num + "\n") // Write in the file
		if err != nil {
			log.Fatal(err)
		}
	}

	f, err = os.Open("data.txt") // Open the file that we create
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scan := bufio.NewScanner(f) // Turn the content of the file in bufio
	for scan.Scan() {           // Scan the file
		fmt.Println("Prime: ", scan.Text()) // Print line per line of the file
	}

	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}
}
