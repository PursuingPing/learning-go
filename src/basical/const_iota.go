package main

import "fmt"

const (
	//可以在const() 添加一个关键字 iota，累加器
	BEIJING = 10 * iota
	SHANGHAI
	SHENZHEN
)

const (
	a, b = iota + 1, iota + 2 //iota = 0, a = 1, b = 2
	c, d                      //iota = 1, c = iota + 1, d = iota + 2
	e, f
	g, h = iota * 2, iota * 3 // iota = 3, g = 6, h = 9
	i, k                      // iota = 4, i = iota * 2 = 8, k = iota *  3 = 12
)

func main() {

	const length int = 10

	fmt.Println("length = ", length)

	fmt.Println("BEIJING = ", BEIJING)
	fmt.Println("SHANGHAI = ", SHANGHAI)
	fmt.Println("SHENZHEN = ", SHENZHEN)

	fmt.Println("a = ", a, "b = ", b)
	fmt.Println("c = ", c, "d = ", d)
	fmt.Println("e = ", e, "f = ", f)
	fmt.Println("g = ", g, "h = ", h)
	fmt.Println("i = ", i, "k = ", k)

	//iota 只能配合const()一起使用，iota只有在const才有累加效果

	println(f2(1))

}

func f2(x int) (_, __ int) {
	_, __ = x, x
	return
}
