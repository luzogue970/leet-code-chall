package main

import (
	"fmt"
	"reflect"
)

func main() {

	// nums := []int{4, 3, 2, 1}
	// excepted := []int{0, 0, 1, 1}
	nums := []int{1, 5, 1, 4, 2}
	excepted := []int{0, 0, 1, 1, 1}

	result := transformArray(nums)
	if reflect.DeepEqual(result, excepted) {
		fmt.Println("yes")
	} else {
		fmt.Printf("nop : %v", result)
	}
}

func transformArray(nums []int) (orderedbianries []int) {

	for _, v := range nums {
		if v%2 == 0 {
			orderedbianries = append([]int{0}, orderedbianries...)
		} else {
			orderedbianries = append(orderedbianries, 1)
		}
	}
	return
}
