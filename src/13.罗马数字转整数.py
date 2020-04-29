#
# @lc app=leetcode.cn id=13 lang=python3
#
# [13] 罗马数字转整数
#

# @lc code=start
class Solution:
    def romanToInt(self, s: str) -> int:
        mp1 = {'I':1, 'V':5, 'X':10, 'L':50, 'C':100, 'D':500, 'M':1000}
        mp2 = {'IV':4, 'IX':9, 'XL':40, 'XC':90, 'CD':400, 'CM':900}
        res = 0
        for k, v in mp2.items():
            if s.find(k) != -1:
                s = s.replace(k, '')
                res += v
        for k, v in mp1.items():
            while s.find(k) != -1:
                s = s.replace(k, '', 1)
                res += v
        return res

# @lc code=end

s = Solution()
print(s.romanToInt('III'))
print(s.romanToInt('IV'))
print(s.romanToInt('IX'))
print(s.romanToInt('LVIII'))
print(s.romanToInt('MCMXCIV'))

