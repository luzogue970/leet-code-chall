package main

import (
	"fmt"
	"reflect"
	"slices"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func createList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}
	head := &ListNode{Val: values[0]}
	current := head
	for _, v := range values[1:] {
		current.Next = &ListNode{Val: v}
		current = current.Next
	}
	return head
}

func main() {

	// values := []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
	// expectedValue := []int{6, 6, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
	// v1 := []int{9, 9, 9, 9, 9, 9, 9}
	// v2 := []int{9, 9, 9, 9}
	// expectedValue := []int{8, 9, 9, 9, 0, 0, 0, 1}
	// expectedValue := []int{6, 6, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}

	v1 := []int{3, 2, 0, 2, 6, 0, 8, 8, 0, 1}
	v2 := []int{0, 8, 9, 6, 8, 7, 2}
	expectedValue := []int{3, 0, 0, 9, 4, 8, 0, 9, 0, 1}

	input1 := createList(v1)
	input2 := createList(v2)

	excepted := createList(expectedValue)
	// input1 := &ListNode{
	// 	Val: 2,
	// 	Next: &ListNode{
	// 		Val: 4,
	// 		Next: &ListNode{
	// 			Val: 3,
	// 		},
	// 	},
	// }
	// input2 := &ListNode{
	// 	Val: 5,
	// 	Next: &ListNode{
	// 		Val: 6,
	// 		Next: &ListNode{
	// 			Val: 4,
	// 		},
	// 	},
	// }

	// excepted := &ListNode{
	// 	Val: 7,
	// 	Next: &ListNode{
	// 		Val: 0,
	// 		Next: &ListNode{
	// 			Val: 8,
	// 		},
	// 	},
	// }
	r := addTwoNumbers(input1, input2)

	if reflect.DeepEqual(r, excepted) {
		println("yes")
	} else {
		printList(r)
	}

}
func getReversedArray(l *ListNode) (a []int) {
	for {
		if l.Next == nil {
			a = append(a, l.Val)
			break
		}
		a = append(a, l.Val)
		l = l.Next
	}
	return
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var r *ListNode
	t1 := getReversedArray(l1)
	t2 := getReversedArray(l2)
	var rArray []int
	lower := min(len(t1), len(t2))
	var retenue int
	for i := 0; i < lower; i++ {
		add := t1[i] + t2[i] + retenue
		retenue = 0
		if add >= 10 {
			retenue = 1
		}
		rArray = append(rArray, add%10)
	}
	var restToAdd []int
	if len(t1) > len(t2) {
		restToAdd = t1
	} else if len(t2) > len(t1) {
		restToAdd = t2
	}
	var r2 []int
	for i := lower; i < len(restToAdd); i++ {
		if restToAdd == nil {
			break
		}
		add := restToAdd[i] + retenue
		retenue = 0
		if add >= 10 {
			retenue = 1
		}
		r2 = append(r2, add%10)
	}
	if retenue == 1 {
		r2 = append(r2, retenue)
	}
	slices.Reverse(r2)
	slices.Reverse(rArray)
	rArray = append(r2, rArray...)
	for _, nb := range rArray {
		newNode := &ListNode{
			Val:  nb,
			Next: r,
		}
		r = newNode
	}
	return r
}

func addTwoNumberswrongToo(l1 *ListNode, l2 *ListNode) *ListNode {
	var rArray []int
	var retenue int
	var r *ListNode

	for {
		addition := l1.Val + l2.Val + retenue
		retenue = 0
		if addition >= 10 {
			rArray = append(rArray, addition%10)
			retenue = 1
		} else {
			rArray = append(rArray, addition)
		}

		if l1.Next == nil || l2.Next == nil {
			break
		}

		l1 = l1.Next
		l2 = l2.Next
	}

	slices.Reverse(rArray)
	for _, nb := range rArray {
		newNode := &ListNode{
			Val:  nb,
			Next: r,
		}
		r = newNode
	}

	var restToAdd *ListNode
	if l1.Next != nil || l2.Next != nil {

		if l1.Next != nil {
			restToAdd = l1
		} else if l2.Next != nil {
			restToAdd = l2
		}
		for {
			if restToAdd.Next == nil || restToAdd == nil {
				break
			}
			addition := restToAdd.Val + retenue
			retenue = 0
			if addition == 10 {
				retenue = 1
				addition = addition % 10
			}
			newNode := &ListNode{
				Val:  addition,
				Next: r,
			}
			r = newNode
			restToAdd = restToAdd.Next
		}
	}
	if retenue == 1 {
		newNode := &ListNode{
			Val:  retenue,
			Next: r,
		}
		r = newNode
	}
	return r
}

func addTwoNumbersBadd(l1 *ListNode, l2 *ListNode) *ListNode {

	nb1 := getReverseNumber(l1)
	nb2 := getReverseNumber(l2)
	intResult := nb1 + nb2
	strRes := strconv.FormatInt(intResult, 10)

	var tInt []int
	for _, runeNb := range strRes {
		intNb, _ := strconv.Atoi(string(runeNb))
		tInt = append(tInt, intNb)
	}

	var r *ListNode

	for _, nb := range tInt {
		newNode := &ListNode{
			Val:  nb,
			Next: r,
		}
		r = newNode
	}
	return r

}

func getReverseNumber(l *ListNode) (reversed int64) {
	var t []string
	var strResult string
	for {
		if l.Next == nil {
			t = append(t, strconv.Itoa(l.Val))
			break
		}
		t = append(t, strconv.Itoa(l.Val))
		l = l.Next

	}
	slices.Reverse(t)
	for _, strNb := range t {
		strResult += strNb
	}
	reversed, _ = strconv.ParseInt(strResult, 10, 64)
	return
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val)
		if head.Next != nil {
			fmt.Print(" -> ")
		}
		head = head.Next
	}
	fmt.Println()
}
