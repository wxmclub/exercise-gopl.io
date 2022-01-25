package main

import (
	"fmt"
	"os"
	"strconv"
)

// 练习 1.2： 修改 echo 程序，使其打印每个参数的索引和值，每个一行。
func main() {
	echo1()
	fmt.Println("------------------------------")
	echo2()
}

func echo1() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + strconv.Itoa(i) + ":" + os.Args[i]
		sep = "\n"
	}
	fmt.Println(s)
}

func echo2() {
	s, sep := "", ""
	for i, arg := range os.Args[1:] {
		s += sep + strconv.Itoa(i+1) + ":" + arg
		sep = "\n"
	}
	fmt.Println(s)
}
