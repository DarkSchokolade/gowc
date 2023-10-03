package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	count_bytes := flag.Bool("c", false, "Byte count")
	flag.Parse()
	files := flag.Args()

	fmt.Println(files)

	var byte_count int64
	for _, file := range files {
		if *count_bytes {
			byte_count = BytesInAFile(file)
		}
		fmt.Printf("%d %s", byte_count, file)
	}
}

// Returns number of bytes
func BytesInAFile(filename string) int64 {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	number_of_bytes := len(data)
	return int64(number_of_bytes)
}
