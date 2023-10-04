package main

import (
	"bufio"
	"flag"
	"fmt"

	"github.com/DarkSchokolade/gowc/wc"
)

func main() {
	count_lines := flag.Bool("l", false, "Print Line count")
	count_words := flag.Bool("w", false, "Print Word count")
	count_bytes := flag.Bool("c", false, "Print Byte count")
	count_chars := flag.Bool("m", false, "Print Character count")
	flag.Parse()
	files := flag.Args()

	for _, file := range files {
		if *count_lines {
			line_count := wc.LinesInAFile(file)
			fmt.Printf("%d\t", line_count)
		}
		if *count_words {
			word_count := wc.Count(file, bufio.ScanWords)
			fmt.Printf("%d\t", word_count)
		}
		if *count_bytes {
			byte_count := wc.BytesInAFile(file)
			fmt.Printf("%d\t", byte_count)
		}
		if *count_chars {
			char_count := wc.Count(file, bufio.ScanRunes)
			fmt.Printf("%d\t", char_count)
		}
		if flag.NFlag() == 0 {
			line_count := wc.LinesInAFile(file)
			fmt.Printf("%d\t", line_count)
			word_count := wc.Count(file, bufio.ScanWords)
			fmt.Printf("%d\t", word_count)
			byte_count := wc.BytesInAFile(file)
			fmt.Printf("%d\t", byte_count)
		}
		fmt.Println(file)
	}
}
