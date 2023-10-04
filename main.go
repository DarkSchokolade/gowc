package main

import (
	"flag"
	"fmt"

	"github.com/DarkSchokolade/gowc/wc"
)

func main() {
	count_bytes := flag.Bool("c", false, "Print Byte count")
	count_lines := flag.Bool("l", false, "Print Line count")
	count_words := flag.Bool("w", false, "Print Word count")
	flag.Parse()
	files := flag.Args()

	for _, file := range files {
		if *count_lines {
			line_count := wc.LinesInAFile(file)
			fmt.Printf("%d\t", line_count)
		}
		if *count_words {
			word_count := wc.WordCountInAFile(file)
			fmt.Printf("%d\t", word_count)
		}
		if *count_bytes {
			byte_count := wc.BytesInAFile(file)
			fmt.Printf("%d\t", byte_count)
		}
		fmt.Println(file)
	}
}
