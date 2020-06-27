package main

/*
 * @lc app=leetcode.cn id=10 lang=golang
 *
 * [10] 正则表达式匹配
 */
// string|dynamic-programming|backtracking

import "fmt"

// 递归截止条件：
// 1. 字符串字面量相等（ s == p 或 len(s) == len(p) == 1 && p[0] == '.'），返回true
// 2. 正则串为空但待匹配串不为空（p == "" && s != ""），返回false

// 分类讨论：
// 场景一、p第二个字符为'*':
//		1、若p第一个字符为'.'，则递归判断s的所有子串（s[0:]...s[len(s):]）和p[2:]
//		2、若p第一个字符为普通字符：
//			例如，bbbcd，b*cd则需要匹配去除b*的s子串与剩余的正则串p[2:]，递归判断bbbcd和cd，bbcd和cd，bcd和cd，cd和cd
//			例如，bbbcd，c*cd则需要匹配去除c*的s子串与剩余的正则串p[2:]，递归判断bbbcd和cd
//		上述场景一，只要有一种情况匹配上，就返回true。
// 场景二、正则串第二个字符不为'*':
//			若两个串的第一个字符相等（s[0] == p[0] || p[0] == '.'），递归判断s[1:]和p[1:]

// @lc code=start
func isMatch(s string, p string) bool {
	if s == p { // 字符串字面量相等
		return true
	} else if len(p) == 1 {
		return len(s) == 1 && p[0] == '.' // 任意单个字符的串和'.'也算匹配
	} else if len(p) == 0 {
		return false
	}

	if p[1] == '*' { // 正则串第二个字符为'*'
		flag := isMatch(s, p[2:])
		for i := 0; i < len(s) && (s[i] == p[0] || p[0] == '.'); i++ {
			flag = flag || isMatch(s[i+1:], p[2:])
		}
		return flag
	} else if len(s) > 0 && (s[0] == p[0] || p[0] == '.') { // 正则串第二个字符不为'*'
		return isMatch(s[1:], p[1:])
	}

	return false
}

// @lc code=end

func main() {
	fmt.Println(isMatch("aa", "a"))
	fmt.Println(isMatch("aa", "a*"))
	fmt.Println(isMatch("aa", ".*"))
	fmt.Println(isMatch("aab", "c*a*b"))
	fmt.Println(isMatch("mississippi", "mis*is*p*."))
	fmt.Println(isMatch("mississippi", "mis*is*ip*."))
}
