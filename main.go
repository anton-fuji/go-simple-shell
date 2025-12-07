package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		if err := execInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func execInput(input string) error {
	input = strings.TrimSpace(input)

	if input == "" {
		return nil
	}

	args := strings.Split(input, " ")
	command := args[0]

	switch command {
	case "cat":
		return cmdCat(args[1:])
	case "ls":
		return cmdLs(args[1:])
	case "grep":
		return cmdGrep(args[1:])
	case "wc":
		return cmdWc(args[1:])
	case "help":
		cmdHelp()
	case "exit":
		fmt.Println("Bye!!")
		os.Exit(0)
	default:
		return fmt.Errorf("command not found %s", command)
	}

	return nil
}

func cmdCat(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("non file oparand")
	}

	for _, filename := range args {
		file, err := os.Open(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cat: %s %v\n", filename, err)
			continue
		}

		_, err = io.Copy(os.Stdout, file)
		file.Close()
		if err != nil {
			return fmt.Errorf("cat: %s %v", filename, err)
		}
	}

	return nil
}

func cmdLs(args []string) error {
	dir := "."
	if len(args) > 0 {
		dir = args[0]
	}

	entries, err := os.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("ls : %v", err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			fmt.Printf("%s/\n", entry.Name())
		} else {
			fmt.Println(entry.Name())
		}
	}

	return nil
}

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

func cmdHelp() {
	title := "MY-SHELL(1)"
	width := 30
	padding := strings.Repeat(" ", width-len(title))

	fmt.Printf("%s%sGeneral Commands Manual%s%s\n\n", title, padding, padding, title)

	fmt.Println("NAME")
	fmt.Println("     my-shell — simple Unix-like shell implemented in Go")
	fmt.Println()

	fmt.Println("SYNOPSIS")
	fmt.Println("     cat FILE [FILE...]")
	fmt.Println("     ls [DIR or FILE]")
	fmt.Println("     grep PATTERN FILE")
	fmt.Println("     wc FILE")
	fmt.Println("     help")
	fmt.Println("     exit")
	fmt.Println()

	fmt.Println("DESCRIPTION")
	fmt.Println("     These commands provide basic file inspection and text processing.")
}
