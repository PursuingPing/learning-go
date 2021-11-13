package main

import "fmt"

func foo1(a string, b int) int {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	c := 100
	return c
}

//多返回值，匿名返回值
func foo2(a string, b int) (int, int) {

	return 666, 888
}

//多返回值，命名返回值
func foo3(a string, b int) (r1 int, r2 int) {
	r1 = 19999
	r2 = 3323223
	return
}

//多返回值，命名返回值
func foo4(a string, b int) (r1, r2 int) {
	fmt.Println("r1 = ", r1)
	fmt.Println("r2 = ", r2)
	r1 = 76357
	r2 = 7546345
	return
}

func main() {
	c := foo1("abc", 8880)
	fmt.Println("c = ", c)

	ret1, ret2 := foo2("hahha", 88809)
	fmt.Println("ret1 = ", ret1, "ret2  = ", ret2)

	ret1, ret2 = foo3("hahha", 88809)
	fmt.Println("ret1 = ", ret1, "ret2  = ", ret2)

	ret1, ret2 = foo4("hahha", 88809)
	fmt.Println("ret1 = ", ret1, "ret2  = ", ret2)
}
