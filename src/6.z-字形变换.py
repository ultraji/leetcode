#
# @lc app=leetcode.cn id=6 lang=python3
#
# [6] Z 字形变换
#

"""
设：步长为 t = (numRows-1)*2
    行索引为 i
第0行字符都是                   0, 1*t, 2*t, k*t
第1 ~ (numRoes-2)行字符都是     0+i, 1*t+i, ..., k*t+i 和 0+t-i, 1*t+t-i, 2*t+t-i, ..., k*t+t-i
第(numRoes-1)行字符都是         0+i, 1*t+i, 2*t+i, ..., k*t+i
"""

# @lc code=start
class Solution:
    def convert(self, s: str, numRows: int) -> str:
        if numRows < 2:
            return s
        res = ""
        slen, t = len(s), (numRows - 1) * 2
        for i in range(numRows):
            for j in range(0, slen, t):
                if i+j >= slen:
                    continue
                res += s[i+j]
                if i != 0 and i != numRows-1 and j+t-i < slen:
                    res += s[j+t-i]
        return res

# @lc code=end
s = Solution()
print(s.convert("LEETCODEISHIRING", 1))
print(s.convert("LEETCODEISHIRING", 2))
print(s.convert("LEETCODEISHIRING", 3))
print(s.convert("LEETCODEISHIRING", 4))
