package main

import (
	"fmt"
)

func twoSum1(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func twoSum2(nums []int, target int) []int {
	// 用 map 存储值到下标的映射
	m := make(map[int]int)
	for i, num := range nums {
		// 判断是否存在 target-num
		if j, ok := m[target-num]; ok {
			return []int{j, i}
		}
		m[num] = i
	}
	return nil
}

/*
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
你可以假设每种输入只会对应一个答案，并且你不能使用两次相同的元素。
你可以按任意顺序返回答案。
*/
// 思路：因为限定了只有一个答案，所以可以暴力遍历
func main() {
	nums := []int{20, 7, 11, 15}
	target := 35
	fmt.Println(twoSum2(nums, target))
	fmt.Println(twoSum1(nums, target))
}
