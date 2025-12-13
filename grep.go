package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cmdGrep(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("grep usage : grep PATTERN FILE")
	}

	pattern := args[0]
	filename := args[1]

	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("grep : %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNum := 0

	for scanner.Scan() {
		lineNum++
		line := scanner.Text()
		if strings.Contains(line, pattern) {
			fmt.Printf("%d:%s\n", lineNum, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("grep: %v", err)
	}
	return nil
}
