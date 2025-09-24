package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		for len(prefix) > 0 && len(strs[i]) >= len(prefix) && strs[i][:len(prefix)] != prefix {
			prefix = prefix[:len(prefix)-1]
		}
		if prefix == "" {
			return ""
		}
	}
	return prefix
}

/*
编写一个函数来查找字符串数组中的最长公共前缀
如果不存在公共前缀，返回空字符串 ""
*/
// 思路：因为是公共前缀，直接默认取第一个字符串为公共前缀
// 然后遍历后面的字符串，如果第一个字符串的前缀与后面的字符串的前缀不相等，则截取公共前缀的最后一个字符，直到相等
func main() {
	strs := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs))
	strs = []string{"dog", "racecar", "car"}
	fmt.Println(longestCommonPrefix(strs))
}
