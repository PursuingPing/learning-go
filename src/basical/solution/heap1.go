package main

import "fmt"

func main() {
	var nums1 []interface{}
	nums2 := []int{1, 2, 3}
	nums3 := append(nums1, nums2)
	fmt.Println(len(nums3))
	fmt.Printf("%v\n", nums3[0])
}
