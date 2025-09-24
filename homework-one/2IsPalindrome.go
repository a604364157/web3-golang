package main

import "strconv"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	s := strconv.Itoa(x)
	reversed := ""
	for i := len(s) - 1; i >= 0; i-- {
		reversed += string(s[i])
	}
	return s == reversed
}

// 给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false
// 解题思路：直接转字符串，然后反转字符串，比较是否相等即可
// 当然也有数学计算方法，不过没啥必要
func main() {
	x := 121
	println(isPalindrome(x))
	x = -121
	println(isPalindrome(x))
	x = 10
	println(isPalindrome(x))
}
