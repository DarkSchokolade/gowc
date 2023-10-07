package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/DarkSchokolade/gowc/wc"
)

var (
	countLines    bool
	countWords    bool
	countBytes    bool
	countChars    bool
	filesToParse  []string
	standardInput *bufio.Scanner
)

func init() {
	flag.BoolVar(&countLines, "l", false, "Print Line count")
	flag.BoolVar(&countWords, "w", false, "Print Word count")
	flag.BoolVar(&countBytes, "c", false, "Print Byte count")
	flag.BoolVar(&countChars, "m", false, "Print Character count")
	flag.Parse()
	filesToParse = flag.Args()
	if flag.NFlag() == 0 {
		flag.Set("l", "true")
		flag.Set("w", "true")
		flag.Set("c", "true")
	}
	if len(filesToParse) == 0 {
		standardInput = bufio.NewScanner(os.Stdin)
	}
}

func main() {
	if standardInput == nil {
		for _, file := range filesToParse {
			fs := ScanFile(file)
			countData := wc.CountAll(*fs)
			PrintCounts(countData)
			fmt.Println(file)
		}
	} else {
		countData := wc.CountAll(*standardInput)
		PrintCounts(countData)
	}
}

func PrintCounts(res wc.ScanResult) {
	if countLines {
		fmt.Printf("%d\t", res.LineCount)
	}
	if countWords {
		fmt.Printf("%d\t", res.WordCount)
	}
	if countBytes {
		fmt.Printf("%d\t", res.ByteCount)
	}
	if countChars {
		fmt.Printf("%d\t", res.CharCount)
	}
}

func ScanFile(filepath string) *bufio.Scanner {
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	// defer f.Close()

	scanner := bufio.NewScanner(f)
	return scanner
}

// func main() {
// 	// sc := bufio.NewScanner(os.Stdin)
// 	// for sc.Scan() {
// 	// 	fmt.Print(sc.Text())
// 	// }
// 	count_lines := flag.Bool("l", false, "Print Line count")
// 	count_words := flag.Bool("w", false, "Print Word count")
// 	count_bytes := flag.Bool("c", false, "Print Byte count")
// 	count_chars := flag.Bool("m", false, "Print Character count")
// 	flag.Parse()
// 	files := flag.Args()

// 	for _, file := range files {
// 		if *count_lines {
// 			line_count := wc.LinesInAFile(file)
// 			fmt.Printf("%d\t", line_count)
// 		}
// 		if *count_words {
// 			scanner := ScanFile(file)
// 			word_count := wc.Count(*scanner, bufio.ScanWords)
// 			fmt.Printf("%d\t", word_count)
// 		}
// 		if *count_bytes {
// 			byte_count := wc.BytesInAFile(file)
// 			fmt.Printf("%d\t", byte_count)
// 		}
// 		if *count_chars {
// 			scanner := ScanFile(file)
// 			char_count := wc.Count(*scanner, bufio.ScanRunes)
// 			fmt.Printf("%d\t", char_count)
// 		}
// 		if flag.NFlag() == 0 {
// 			scanner := ScanFile(file)
// 			line_count := wc.LinesInAFile(file)
// 			fmt.Printf("%d\t", line_count)
// 			word_count := wc.Count(*scanner, bufio.ScanWords)
// 			fmt.Printf("%d\t", word_count)
// 			byte_count := wc.BytesInAFile(file)
// 			fmt.Printf("%d\t", byte_count)
// 		}
// 		if len(files) == 0 {
// 			sc := bufio.NewScanner(os.Stdin)
// 			// call the respective flag method
// 		}
// 		fmt.Println(file)
// 	}
// }
