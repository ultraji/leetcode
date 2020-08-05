package main

/*
 * @lc app=leetcode.cn id=24 lang=golang
 *
 * [24] 两两交换链表中的节点
 */
// linked-list

import "fmt"

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func swapPairs(head *ListNode) *ListNode {
	ans := &ListNode{0, head}
	for p1, p2 := ans, ans.Next; p1 != nil && p2 != nil && p2.Next != nil ; p1, p2 = p2, p2.Next {
		t := p2.Next
		p2.Next = t.Next
		t.Next = p2
		p1.Next = t
	}
	return ans.Next
}

// @lc code=end
func main() {
	fmt.Println(swapPairs(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}))
}
