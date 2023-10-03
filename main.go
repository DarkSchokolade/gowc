package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	filename := flag.String("c", "test.txt", "Byte count")
	// f := flag.Arg(0)
	flag.Parse()
	f := flag.Args()

	fmt.Println(f)

	data, err := os.ReadFile(*filename)
	if err != nil {
		log.Fatal(err)
	}
	number_of_bytes := len(data)
	fmt.Printf("%d %s", number_of_bytes, *filename)
}
