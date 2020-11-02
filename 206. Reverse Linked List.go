package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var tmp *ListNode
	now := head

	for {
		if now == nil {
			break
		}
		i := now
		now = now.Next
		i.Next = tmp
		tmp = i
	}

	for {
		if tmp == nil {
			break
		}
		fmt.Println(tmp.Val)
		tmp = tmp.Next
	}
	return tmp
}
