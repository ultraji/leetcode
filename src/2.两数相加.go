package main

/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	c, p := 0, res
	for ; l1 != nil && l2 != nil; l1, l2 = l1.Next, l2.Next {
		t := l1.Val + l2.Val + c
		p.Next, c = &ListNode{t % 10, nil}, t/10
		p = p.Next
	}
	l := l1
	if l == nil {
		l = l2
	}
	for ; l != nil; l = l.Next {
		t := l.Val + c
		p.Next, c = &ListNode{t % 10, nil}, t/10
		p = p.Next
	}
	if c != 0 {
		p.Next = &ListNode{c, nil}
		p = p.Next
	}
	return res.Next
}

// @lc code=end

// func main() {
// 	l1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
// 	l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
// 	addTwoNumbers(l1, l2)
// }
