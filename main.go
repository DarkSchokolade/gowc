package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	count_bytes := flag.Bool("c", false, "Byte count")
	count_lines := flag.Bool("l", false, "Line count")
	flag.Parse()
	files := flag.Args()

	fmt.Println(files)

	var byte_count int64
	var line_count int
	for _, file := range files {
		if *count_bytes {
			byte_count = BytesInAFile(file)
			fmt.Printf("%d\t", byte_count)
		}
		if *count_lines {
			line_count = LinesInAFile(file)
			fmt.Printf("%d\t", line_count)
		}
		fmt.Println(file)
	}
}

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
