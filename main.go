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

	// ディレクトリないのファイルを一覧取得j
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
