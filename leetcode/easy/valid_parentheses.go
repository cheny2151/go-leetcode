package easy

// 有效的括号 https://leetcode.cn/problems/valid-parentheses/description/
// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
// 有效字符串需满足：
// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
// 每个右括号都有一个对应的相同类型的左括号。
//
// 示例 1：
// 输入：s = "()"
// 输出：true
//
// 示例 2：
// 输入：s = "()[]{}"
// 输出：true
//
// 示例 3：
// 输入：s = "(]"
// 输出：false

type node struct {
	r   rune
	pre *node
}

func isValid(s string) bool {
	runes := []rune(s)
	var cur *node
	for i := 0; i < len(runes); i++ {
		r := runes[i]
		if r == '(' || r == '[' || r == '{' {
			cur = &node{
				r:   r,
				pre: cur,
			}
			continue
		}
		if cur == nil {
			return false
		}
		pre := cur.r
		switch r {
		case ')':
			if pre != '(' {
				return false
			}
		case ']':
			if pre != '[' {
				return false
			}
		case '}':
			if pre != '{' {
				return false
			}
		}
		cur = cur.pre
	}
	return cur == nil
}
