package main

import (
	"fmt"
	//匿名导包，可以不使用
	_ "learning-go/src/basical/init/lib1"
	mylib2 "learning-go/src/basical/init/lib2"
	//少用.
	//. "learning-go/src/basical/init/lib2"
)

func main() {
	fmt.Println("main is running")
	//lib1.Lib1Test()
	mylib2.Lib2Test()
}

func init() {
	fmt.Println("init.go inti() is running")
}
