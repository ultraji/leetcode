package main

/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第N个节点
 */
// linked-list|two-pointers

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	newHead := &ListNode{0, head}
	p1, p2 := newHead, newHead
	for ; n > 0 && p2.Next != nil ; n-- {
		p2 = p2.Next
	}
	for p2.Next != nil {
		p1, p2 = p1.Next, p2.Next
	}
	if p1.Next != nil {
		p1.Next = p1.Next.Next
	}
	return newHead.Next
}
// @lc code=end

