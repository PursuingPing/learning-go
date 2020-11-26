package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//1
	//var s, sep string
	//fmt.Println(len(os.Args))
	//for i := 0; i < len(os.Args); i++ {
	//	s += sep + os.Args[i]
	//	sep = " "
	//}
	//2
	//s, sep := "", ""
	//for _, arg := range os.Args[0 :] {
	//	s += sep + arg
	//	sep = " "
	//}
	//fmt.Println(s)

	//3
	fmt.Println(strings.Join(os.Args[0:], " "))
}
