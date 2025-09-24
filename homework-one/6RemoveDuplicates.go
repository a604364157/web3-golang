package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	slow := 0
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
	}
	return slow + 1
}

/*
给你一个有序数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次，返回删除后数组的新长度。
不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
*/
// 思路：快慢指针
// 定义快指针fast，慢指针slow
// 快指针从第二个元素开始遍历，慢指针从第一个元素开始遍历
// 如果快指针指向的元素和慢指针指向的元素不同，则将快指针指向的元素赋值给慢指针指向的元素，并将慢指针加1
func main() {
	nums := []int{2, 2, 2, 3}
	fmt.Println(removeDuplicates(nums))
}
