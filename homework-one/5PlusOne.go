package main

import "fmt"

func plusOne(digits []int) []int {
	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		digits[i]++
		if digits[i] < 10 {
			return digits
		}
		digits[i] = 0
	}
	result := make([]int, n+1)
	result[0] = 1
	return result
}

/*
 给定一个表示 大整数 的整数数组 digits，其中 digits[i] 是整数的第 i 位数字。
 这些数字按从左到右，从最高位到最低位排列。这个大整数不包含任何前导 0。
 将大整数加 1，并返回结果的数字数组。
*/
// 思路：
// 1:直接转换为数字，然后再加一，缺陷是可能会溢出(故不用此方法)
// 2:从最低位开始，如果当前位数字不为 9，则加一，并返回
// 3:如果当前位数字为 9，则置为 0，继续往下一位
func main() {
	digits := []int{1, 2, 3}
	fmt.Println(plusOne(digits))
	digits = []int{4, 3, 2, 1}
	fmt.Println(plusOne(digits))
	digits = []int{9}
	fmt.Println(plusOne(digits))
	digits = []int{9, 9, 9, 9, 9, 9, 9}
	fmt.Println(plusOne(digits))
}
