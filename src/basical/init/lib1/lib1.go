package lib1

import "fmt"

func Lib1Test() {
	fmt.Println("lib1 ")
}

func init() {
	fmt.Println("lib1 init() is running")
}
