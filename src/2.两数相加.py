#
# @lc app=leetcode.cn id=2 lang=python3
#
# [2] 两数相加
#

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
    def addTwoNumbers(self, l1: ListNode, l2: ListNode) -> ListNode:
        c, res = 0, ListNode(0)
        node = res
        while l1 and l2:
            t = c + l1.val + l2.val
            node.next, c = ListNode(t%10), t//10
            l1, l2, node = l1.next, l2.next, node.next
        l = l1 if l1 else l2
        while l:
            t = c + l.val
            node.next, c = ListNode(t%10), t//10
            l, node = l.next, node.next
        if c != 0:
            node.next = ListNode(c)
        return res.next

# @lc code=end
