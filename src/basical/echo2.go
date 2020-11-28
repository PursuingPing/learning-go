package main

import (
	"flag"
	"fmt"
	"strings"
)

//flag  使用程序的命令行参数来设置整个程序内的某些变量值
var n = flag.Bool("n", false, "omit trailing newLine")
var sep = flag.String("s", " ", "separator")

func main() {
	//先更新标识变量的默认值
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
	fmt.Printf("gcd : %d \n", gcd(12, 8))
	z := 0.0
	fmt.Println(z, -z, 1/z, -1/z, z/z)

	var x complex128 = complex(1, 2)
	var y complex128 = complex(3, 4)
	fmt.Println(x * y)
	fmt.Println(real(x * y))
	fmt.Println(imag(x * y))
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}
