package main

import (
	"bufio"
	"fmt"
	"os"
)

// 练习 1.4： 修改 dup2 ，出现重复的行时打印文件名称。
func main() {
	counts := make(map[string]int)
	fileTags := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, fileTags)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, fileTags)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, line, fileTags[line])
		}
	}
}

func countLines(f *os.File, counts map[string]int, fileTags map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		counts[line]++
		// 处理重复文件名
		repeat := false
		for _, fileTag := range fileTags[line] {
			if fileTag == f.Name() {
				repeat = true
			}
		}
		if !repeat {
			fileTags[line] = append(fileTags[line], f.Name())
		}
	}
	// NOTE: ignoring potential errors from input.Err()
}
