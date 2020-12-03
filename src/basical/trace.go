package main

import (
	"fmt"
	"time"
)

func bigSlowOperation() {
	defer trace("bigSlowOperation")()
	fmt.Println("running...")
	time.Sleep(10 * time.Second)
	d := []int{1, 2, 3}
	fmt.Println(d[5])
	fmt.Println("ending...")
}

func trace(msg string) func() {
	start := time.Now()
	fmt.Printf("enter %s\n", msg)
	return func() {
		fmt.Printf("exit %s (%s)\n", msg, time.Since(start))
	}
}

func main() {
	bigSlowOperation()
}
