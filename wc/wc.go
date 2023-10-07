package wc

import (
	"bufio"
	"log"
	"os"
	"strings"
	"unicode/utf8"
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

// Counts number of tokens in a file.
// Tokenization is done based on bufio.SplitFunc type
func Count(scanner bufio.Scanner, split bufio.SplitFunc) int {
	count := 0
	// Set the split function for the scanning operation
	scanner.Split(split)
	for scanner.Scan() {
		count++
	}
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	return count
}

type ScanResult struct {
	LineCount int
	WordCount int
	ByteCount int
	CharCount int
}

func CountAll(scanner bufio.Scanner) ScanResult {
	lineCount, wordCount, byteCount, charCount := 0, 0, 0, 0
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		lineCount++

		charCount += utf8.RuneCountInString(line)

		byteCount += len(line)

		words := strings.Split(line, " ")
		wordCount += len(words)
	}

	return ScanResult{
		LineCount: lineCount,
		WordCount: wordCount,
		ByteCount: byteCount,
		CharCount: charCount,
	}
}
