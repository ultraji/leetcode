package main

/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并K个排序链表
 */
// linked-list|divide-and-conquer|heap

import (
	"container/heap"
	"fmt"
)

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ptrHeap []*ListNode

func (h ptrHeap) Len() int {
	return len(h)
}

func (h ptrHeap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h ptrHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *ptrHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *ptrHeap) Pop() interface{} {
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}

func mergeKLists(lists []*ListNode) *ListNode {
	ans := &ListNode{}
	hp := &ptrHeap{}
	length := 0
	for i := 0; i < len(lists); i++ {
		if lists[i] == nil {
			continue
		}
		*hp = append(*hp, lists[i])
		length++
	}
	heap.Init(hp)
	for p := ans; length != 0; p = p.Next {
		p.Next = (*hp)[0]
		ptr := (*hp)[0].Next
		if ptr != nil {
			(*hp)[0] = ptr
			heap.Fix(hp, 0)
		} else {
			heap.Remove(hp, 0)
			length--
		}
	}
	return ans.Next
}

// @lc code=end
func main() {
	fmt.Println(mergeKLists([]*ListNode{
		&ListNode{1, &ListNode{4, &ListNode{5, nil}}},
		&ListNode{1, &ListNode{3, &ListNode{4, nil}}},
		&ListNode{2, &ListNode{6, nil}},
	}))
}
