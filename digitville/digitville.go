package main

import (
	"fmt"
	"reflect"
)

//https://leetcode.com/problems/the-two-sneaky-numbers-of-digitville/ 
//done

func main() {
	nums := []int{0,3,2,1,3,2}
	wanted := []int{1, 0}
	sneakyNumber := getSneakyNumbers(nums)
	if reflect.DeepEqual(sneakyNumber, wanted) {
		fmt.Println("yes")
	} else {
		fmt.Printf("nop : wanted :  %v, got : %v", wanted, sneakyNumber)
	}
}

func getSneakyNumbers(nums []int)  (sneakyOnes []int) {
	var gottenNumbers []int
	for _, num := range nums {
		for _, gottenNumber := range gottenNumbers {
			if gottenNumber == num {
				sneakyOnes = append(sneakyOnes, num)
				break
			}
		}
		gottenNumbers = append(gottenNumbers, num)
	}
	return
}
