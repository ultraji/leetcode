#
# @lc app=leetcode.cn id=22 lang=python3
#
# [22] 括号生成
# backtracking

from typing import List

# @lc code=start
class Solution:
    def generateParenthesis(self, n: int) -> List[str]:
        g = self.generateParenthesis
        if n == 1:
            return ["()"]
        else:
            res = [ "(" + s + ")" for s in g(n-1)]
            for i in range(1, n):
                res += [x+y for x in g(i) for y in g(n-i)]
            return list(set(res))

# @lc code=end

s = Solution()
print(s.generateParenthesis(3))
print(s.generateParenthesis(4))