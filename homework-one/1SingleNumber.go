package main

func singleNumber(nums []int) int {
	result := 0
	for _, num := range nums {
		result ^= num
	}
	return result
}

// 只出现一次的数字
// 给你一个 非空 整数数组 nums ，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素
// 解题思路：因为只有两种情况，一种是出现一次，一种是出现两次，所以可以用异或运算来实现
func main() {
	nums := []int{2, 2, 1}
	println(singleNumber(nums))
	nums = []int{4, 1, 2, 1, 2}
	println(singleNumber(nums))
	nums = []int{1}
	println(singleNumber(nums))
}
