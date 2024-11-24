package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// 引数のチェック
	if len(os.Args) < 4 {
		fmt.Println("Usage: file_manipulator <command> <inputpath> <outputpath>")
		return
	}

	command := os.Args[1]
	inputpath := os.Args[2]

	// コマンドの実行
	switch command {
	case "reverse":
		if len(os.Args) != 4 {
			fmt.Println("Usage: reverse <inputpath> <outputpath>")
			return
		}
		if err := reverse(inputpath, os.Args[3]); err != nil {
			fmt.Println("Error:", err)
		}
	case "copy":
		if len(os.Args) != 4 {
			fmt.Println("Usage: copy <inputpath> <outputpath>")
			return
		}
		if err := copy(inputpath, os.Args[3]); err != nil {
			fmt.Println("Error:", err)
		}
	case "duplicate-contents":
		if len(os.Args) != 4 {
			fmt.Println("Usage: duplicate-contents <inputpath> <n>")
			return
		}
		if err := duplicate(inputpath, os.Args[3]); err != nil {
			fmt.Println("Error:", err)
		}
	case "replace-string":
		if len(os.Args) != 5 {
			fmt.Println("Usage: replace-string <inputpath> <searchString> <newString>")
			return
		}
		if err := replace(inputpath, os.Args[3], os.Args[4]); err != nil {
			fmt.Println("Error:", err)
		}
	default:
		fmt.Println("Unknown command:", command)
	}

}

// ファイルを逆順にして出力する関数
func reverse(inputpath string, outputpath string) error {
	contents, err := readFile(inputpath)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	reversed := reverseString(contents)

	if err := writeFile(outputpath, reversed); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

func reverseString(s string) string {
	// 文字列を []rune に変換
	// Goの文字列はUTF-8でエンコードされたバイト列として扱われる。
	// []runeに変換することで、文字列を「Unicodeコードポイント」単位（文字単位）で扱えるようになる。
	// これにより、マルチバイト文字（例: 日本語や絵文字）も正しく操作できる。
	runes := []rune(s)

	// スライス内の要素を反転する処理
	// i: 左端のインデックス, j: 右端のインデックス
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		// 左端と右端の要素を交換
		runes[i], runes[j] = runes[j], runes[i]
	}

	// 反転後の []rune を文字列に戻して返す
	// string()に変換することで、[]runeを再び文字列に戻す。
	return string(runes)
}

func copy(inputpath string, outputpath string) error {
	contents, err := readFile(inputpath)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	if err := writeFile(outputpath, contents); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

func duplicate(inputpath string, n string) error {
	contents, err := readFile(inputpath)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	num, err := strconv.Atoi(n)
	if err != nil {
		return fmt.Errorf("failed to convert n to int: %w", err)
	}

	duplicatedContents := contents
	for i := 1; i < num; i++ {
		duplicatedContents += contents
	}

	if err := writeFile(inputpath, duplicatedContents); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}


func replace(inputpath string, searchString string, newString string) error {
	contents, err := readFile(inputpath)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	// searchString を newString に置き換え
	replacedContents := strings.ReplaceAll(contents, searchString, newString)

	if err := writeFile(inputpath, replacedContents); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}


func readFile(filepath string) (string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var contents string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		contents += scanner.Text() + "\n"
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return contents, nil
}


func writeFile(filepath string, content string) error {
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}
