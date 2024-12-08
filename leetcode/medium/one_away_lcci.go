package medium

import "math"

// https://leetcode.cn/problems/one-away-lcci/description/
// 字符串有三种编辑操作:插入一个英文字符、删除一个英文字符或者替换一个英文字符。 给定两个字符串，编写一个函数判定它们是否只需要一次(或者零次)编辑。
//示例 1：
//输入：
//first = "pale"
//second = "ple"
//输出：True

//示例 2：
//输入：
//first = "pales"
//second = "pal"
//输出：False

func oneEditAway(first string, second string) bool {
	fr := []rune(first)
	sr := []rune(second)
	if len(fr) > len(sr)+1 || len(sr) > len(fr)+1 {
		return false
	}
	fix := false
	fi, si := 0, 0
	for fi < len(fr) && si < len(sr) {
		if fr[fi] == sr[si] {
			fi++
			si++
			continue
		}
		if fix {
			return false
		}
		if fi+1 < len(fr) && fr[fi+1] == sr[si] {
			fi++
			fix = true
		} else if si+1 < len(sr) && fr[fi] == sr[si+1] {
			si++
			fix = true
		} else {
			fi++
			si++
			fix = true
		}
	}
	if fix && fi >= len(fr) && si >= len(sr) {
		return true
	} else if !fix && math.Abs(float64(len(fr)-len(sr))) < 2 {
		return true
	}
	return false
}
