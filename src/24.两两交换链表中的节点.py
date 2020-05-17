#
# @lc app=leetcode.cn id=24 lang=python3
#
# [24] 两两交换链表中的节点
#

# @lc code=start
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def swapPairs(self, head: ListNode) -> ListNode:
        new_head = ListNode(0)
        new_head.next = head
        p1, p2 = new_head, new_head.next
        while p1 and p2:
            t = p2.next
            if t:
                m = t.next
                p1.next = t
                t.next = p2
                p2.next = m
            p1, p2 = p2, p2.next
        return new_head.next
# @lc code=end

