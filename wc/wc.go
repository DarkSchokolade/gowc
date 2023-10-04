package wc

import (
	"bufio"
	"log"
	"os"
)

// Returns number of bytes in a file
func BytesInAFile(filename string) int64 {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	number_of_bytes := len(data)
	return int64(number_of_bytes)
}

// Returns number of lines in a file
func LinesInAFile(filepath string) int {
	total_lines := 0
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		total_lines++
	}

	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	return total_lines
}

// Returns number of words in a file.
func WordCountInAFile(filepath string) int {
	count := 0
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	// Set the split function for the scanning operation
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		count++
	}
	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	return count
}
