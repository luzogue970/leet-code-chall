package main

import (
	"fmt"
)

//https://leetcode.com/problems/insert-greatest-common-divisors-in-linked-list/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	input := ListNode{
		Val: 18,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  10,
				Next: &ListNode{Val: 3},
			},
		},
	}
	fmt.Println("Avant :")
	printList(&input)
	output := insertGreatestCommonDivisors(&input)

	fmt.Println("Après :")
	printList(output)
}

func insertGreatestCommonDivisors(head *ListNode)  *ListNode {
	current := head

	for current != nil && current.Next != nil {

		v1 := current.Val
		v2 := current.Next.Val

		minV := min(v1, v2)

		var vUp int
		var vLow int

		if minV != v1 {
			vUp, vLow = v1, v2
		} else {
			vUp, vLow = v2, v1
		}

		for i := minV; i > 0; i-- {
			if vUp%i == 0 && vLow%i == 0 {
				newNode := &ListNode{Val: i, Next: current.Next}
				current.Next = newNode
				current = newNode.Next
				break
			}
		}
	}
	return head
}

// Fonction utilitaire pour afficher une liste chaînée
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
