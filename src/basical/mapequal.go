package main

import "fmt"

func mapequal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}

func main() {
	var nilmap map[string]int
	fmt.Println(mapequal(map[string]int{"A": 0}, map[string]int{"B": 45}))
	fmt.Println(mapequal(map[string]int{"A": 0}, map[string]int{"A": 1}))
	fmt.Println(mapequal(map[string]int{"A": 0}, map[string]int{"A": 0}))
	fmt.Println(mapequal(map[string]int{"A": 0}, nilmap))

}
