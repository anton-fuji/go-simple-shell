package main

import (
	"fmt"
	"io"
	"os"
)

func cmdCat(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("non file operand")
	}

	for _, filename := range args {
		file, err := os.Open(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cat: %s %v\n", filename, err)
			continue
		}
		defer file.Close()
		_, err = io.Copy(os.Stdout, file)
		if err != nil {
			return fmt.Errorf("cat: %s %v", filename, err)
		}
	}

	return nil
}
