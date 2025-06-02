package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	in := "001011"
	excepted := []int{11, 8, 5, 4, 3, 4}
	r := minOperations(in)
	if reflect.DeepEqual(r, excepted) {
		fmt.Println("yes")
	} else {
		fmt.Printf("%v", r)

	}

}

func minOperations(boxes string)  []int {
	n := len(boxes)
	r := make([]int, n)
	for i := 0; i < n; i++ {
		for movingBoxeIndex, letter := range boxes {
			if letter == '1' {
				diff := math.Abs(float64(i) - float64(movingBoxeIndex))
				r[i] += int(diff)
			}
		}
	}
	return r
}
