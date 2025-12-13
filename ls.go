package main

import (
	"fmt"
	"os"
)

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
