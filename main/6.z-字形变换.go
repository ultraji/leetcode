package main

/*
 * @lc app=leetcode.cn id=6 lang=golang
 *
 * [6] Z 字形变换
 */

/*
 * 设：步长为 t = (numRows-1)*2
 * 行索引为 i
 * 第0行字符都是                   0, 1*t, 2*t, k*t
 * 第1 ~ (numRoes-2)行字符都是     0+i, 1*t+i, ..., k*t+i 和 0+t-i, 1*t+t-i, 2*t+t-i, ..., k*t+t-i
 * 第(numRoes-1)行字符都是         0+i, 1*t+i, 2*t+i, ..., k*t+i
 */

// @lc code=start
func convert(s string, numRows int) string {
	if numRows < 2 {
		return s
	}
	var res []byte
	slen, t := len(s), (numRows-1)*2
	for i := 0; i < numRows; i++ {
		for j := 0; j < slen; j += t {
			if i+j >= slen {
				continue
			}
			res = append(res, s[i+j])
			if i != 0 && i != numRows-1 && j+t-i < slen {
				res = append(res, s[j+t-i])
			}
		}
	}
	return string(res)
}

// @lc code=end
// func main() {
// 	fmt.Println(convert("LEETCODEISHIRING", 3))
// }
