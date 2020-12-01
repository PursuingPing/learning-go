package main

import "fmt"

func noempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func noempty2(strings []string) []string {
	out := strings[:0]
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func main() {
	//slice
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", noempty(data))
	fmt.Printf("%q\n", data)

	data2 := []string{"one", "", "three"}
	fmt.Printf("%q\n", noempty2(data2))
	fmt.Printf("%q\n", data2)

	s := []int{5, 6, 7, 8, 9}
	fmt.Println(remove(s, 2))
}
