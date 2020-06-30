package main

/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 */
// linked-list

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	newHead := &ListNode{0, nil}
	p := newHead
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			p.Next, l1 = l1, l1.Next
		} else {
			p.Next, l2 = l2, l2.Next
		}
		p = p.Next
	}
	if l1 != nil {
		p.Next = l1
	}else if l2 != nil {
		p.Next = l2
	}
	return newHead.Next
}
// @lc code=end

