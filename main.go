package main

import (
	"bufio"
	"fmt"
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
