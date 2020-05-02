#
# @lc app=leetcode.cn id=12 lang=python3
#
# [12] 整数转罗马数字
#

# @lc code=start
class Solution:
    def intToRoman(self, num: int) -> str:
        mp = {1:'I', 5:'V', 10:'X', 50:'L', 100:'C', 500:'D', 1000:'M', 
            4:'IV', 9:'IX', 40:'XL', 90:'XC', 400:'CD', 900:'CM'}
        row, roman = [1000, 100, 10 ,1], ''
        for r in row:
            if num >= r:
                t, num = num//r, num%r
                if t in [4, 5, 9]:
                    roman += mp[t*r]
                else:
                    if t > 5:
                        roman += mp[5*r]
                        t -= 5
                    roman += mp[r] * t
        return roman
# @lc code=end

s = Solution()
print(s.intToRoman(3))
print(s.intToRoman(4))
print(s.intToRoman(9))
print(s.intToRoman(58))
print(s.intToRoman(1994))