package main

func isValid(s string) bool {
	var stack []rune
	pairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, ch := range s {
		switch ch {
		case '(', '{', '[':
			stack = append(stack, ch) // 入栈
		case ')', '}', ']':
			if len(stack) == 0 || stack[len(stack)-1] != pairs[ch] {
				return false
			}
			stack = stack[:len(stack)-1] // 出栈
		}
	}
	return len(stack) == 0
}

/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
有效字符串需满足：
左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
每个右括号都有一个对应的相同类型的左括号。
*/
// 思路：
// 1. 遍历字符串，遇到左括号入栈，遇到右括号出栈，如果栈为空或者栈顶元素不是对应的左括号，则返回false
// 2. 遍历完成后，如果栈为空，则返回true，否则返回false
func main() {
	s := "()[]{}"
	println(isValid(s))
	s = "([)]"
	println(isValid(s))
	s = "{[]}"
	println(isValid(s))
	s = "{([])}"
	println(isValid(s))

}
