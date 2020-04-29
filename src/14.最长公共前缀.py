#
# @lc app=leetcode.cn id=14 lang=python3
#
# [14] 最长公共前缀
#

from typing import List

# @lc code=start
class Solution:
    def longestCommonPrefix(self, strs: List[str]) -> str:
        s = ""
        for x in zip(*strs):
            if len(set(x)) != 1:
                break
            s += x[0]
        return s
# @lc code=end

s = Solution()
print(s.longestCommonPrefix(["dog","racecar","car"]))
print(s.longestCommonPrefix(["flower","flow","flight"]))