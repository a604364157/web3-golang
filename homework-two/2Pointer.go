package main

import "fmt"

func multiplyTwo(nums *[]int) {
	for i := range *nums {
		(*nums)[i] *= 2
	}
}

/*
实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2
*/
func main() {
	nums := []int{1, 2, 3, 4, 5}
	multiplyTwo(&nums)
	fmt.Println(nums)
}
