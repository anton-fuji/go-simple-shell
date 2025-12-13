package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cmdWc(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("missing file operand")
	}

	var totalLines, totalWords, totalBytes int
	var hasErr bool

	for _, filename := range args {
		file, err := os.Open(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "wc: %s: %v\n", filename, err)
			continue
		}

		lines, words, bytes, err := countFile(file)
		file.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "wc: %s %v\n", filename, err)
			hasErr = true
			continue
		}

		fmt.Printf("%7d %7d %7d %s\n", lines, words, bytes, filename)

		totalLines += lines
		totalWords += words
		totalBytes += bytes
	}

	if len(args) > 1 {
		fmt.Printf("%7d %7d %7d total\n", totalLines, totalWords, totalBytes)
	}
	if hasErr {
		return fmt.Errorf("wc: some files could not be read")
	}

	return nil
}

func countFile(file *os.File) (lines, words, bytes int, err error) {
	stat, err := file.Stat()
	if err != nil {
		return 0, 0, 0, err
	}
	bytes = int(stat.Size())

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines++
		words += len(strings.Fields(line))
	}

	if err := scanner.Err(); err != nil {
		return 0, 0, 0, fmt.Errorf("read error: %w", err)
	}

	return
}
