package main

import (
	"fmt"
	"os"
	"strings"
)

// 练习 1.1： 修改 echo 程序，使其能够打印 os.Args[0] ，即被执行命令本身的名字。
func main() {
	echo1()
	fmt.Println("------------------------------")
	echo2()
	fmt.Println("------------------------------")
	echo3()
}

func echo1() {
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func echo2() {
	s, sep := "", ""
	for _, arg := range os.Args[0:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func echo3() {
	fmt.Println(strings.Join(os.Args[0:], " "))
}
