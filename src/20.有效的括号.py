#
# @lc app=leetcode.cn id=20 lang=python3
#
# [20] 有效的括号
#

# @lc code=start
class Solution:
    def isValid(self, s: str) -> bool:
        st = list()
        for c in s:
            if c in ['(','[','{']:
                st.append(c)
            elif c in [')',']','}']:
                if not st:
                    return False
                t = st.pop()
                if t+c not in ['()','[]','{}']:
                    return False
        if st:
            return False
        else:
            return True
# @lc code=end

