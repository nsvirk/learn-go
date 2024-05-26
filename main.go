package main

import (
	"fmt"

	"github.com/nsvirk/learn-go/ch01"
	"github.com/nsvirk/learn-go/ch03"
	"github.com/nsvirk/learn-go/ch04"
	"github.com/nsvirk/learn-go/ch05"
)

func init() {
	fmt.Println("main package initialized")
}

func main() {
	// main
	fmt.Println("func Main()")

	// ch01
	ch01.LearnFmt()
	// ch01.LearnFmtPrintf()
	// ch01.LearnFmtSprintf()
	fmt.Println("------------------")

	// ch02
	// ch02.PrintDataTypes()
	// ch02.PrintPersonStruct()
	// ch02.MarshalJson()
	fmt.Println("------------------")

	// ch03
	ch03.LearnControlStructures()
	fmt.Println("------------------")

	// ch04
	ch04.LearnFunctions()
	fmt.Println("------------------")

	// ch05
	ch05.LearnPointers()

}
