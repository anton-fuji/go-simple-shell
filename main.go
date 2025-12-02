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

		// ここでcat が入力されてコマンドが実行される処理を書く
		if err := cmdInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func cmdInput(input string) error {
	// 改行を削除して前後の空白を無くしていく
	input = strings.TrimSpace(input)

	if input == "" {
		return nil
	}

	// スペースで分割して、コマンドと引数を分ける
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

	// 各ファイルの処理を順に行う
	for _, filename := range args {
		file, err := os.Open(filename)
		if err != nil {
			// 標準エラーを出力
			fmt.Fprintf(os.Stderr, "cat: %s %v\n", filename, err)
			continue
		}

		// ファイルの内容を標準出力にコピー
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

	// ディレクトリないのファイルを一覧取得
	entries, err := os.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("ls : %v", err)
	}

	// ファイル名を表示していく
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

	// 行ごとにスキャン
	for scanner.Scan() {
		lineNum++
		line := scanner.Text()
		//パターンが含まれていれば表示
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

	filename := args[0]

	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("wc: %v", err)
	}
	defer file.Close()

	var (
		lines int
		words int
		bytes int
	)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines++
		line := scanner.Text()
		bytes += len(line) + 1 // +1 for newline
		words += len(strings.Fields(line))
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("wc: %v", err)
	}

	fmt.Printf("%d %d %d %s\n", lines, words, bytes, filename)

	return nil
}
