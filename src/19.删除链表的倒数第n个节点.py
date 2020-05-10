#
# @lc app=leetcode.cn id=19 lang=python3
#
# [19] 删除链表的倒数第N个节点
#  two-pointers

class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

# @lc code=start
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def removeNthFromEnd(self, head: ListNode, n: int) -> ListNode:
        new_head = ListNode(0)
        new_head.next = head
        p1, p2 = new_head, new_head
        while n > 0 and p2.next:
            p2 = p2.next
            n -= 1
        while p2.next:
            p1, p2 = p1.next, p2.next
        if p1.next:
            p1.next = p1.next.next
        return new_head.next
# @lc code=end

