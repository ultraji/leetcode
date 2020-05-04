#
# @lc app=leetcode.cn id=17 lang=python3
#
# [17] 电话号码的字母组合
#

from typing import List

# @lc code=start
class Solution:
    def letterCombinations(self, digits: str) -> List[str]:
        if not digits:
            return []
        mp = ['', '', 'abc', 'def', 'ghi', 'jkl', 'mno', 'pqrs', 'tuv', 'wxyz']
        res = ['']
        for d in digits:
            table = mp[int(d)]
            res = [x+y for x in res for y in table]
        return res
# @lc code=end

s = Solution()
print(s.letterCombinations('23'))