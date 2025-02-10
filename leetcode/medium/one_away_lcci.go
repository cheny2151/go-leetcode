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
	// 双指针，从头遍历，直到至少有一个完成遍历
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
	if fix && fi == len(fr) && si == len(sr) {
		// 1：fix过的则必须都完成了遍历
		return true
	} else if !fix && math.Abs(float64(len(fr)-len(sr))) < 2 {
		// 2：未fix过的则长度最多只能相差1
		return true
	}
	return false
}

// 编辑距离解法：性能较差
func oneEditAway2(first string, second string) bool {
	fr := []rune(first)
	sr := []rune(second)

	dp := make([][]int, len(fr)+1)
	for i := 0; i <= len(fr); i++ {
		dp[i] = make([]int, len(sr)+1)
		dp[i][0] = i
	}
	for i := 0; i <= len(sr); i++ {
		dp[0][i] = i
	}

	for i := 1; i <= len(fr); i++ {
		for j := 1; j <= len(sr); j++ {
			plan := min(dp[i-1][j]+1, dp[i][j-1]+1)
			plan2 := dp[i-1][j-1]
			if fr[i-1] != sr[j-1] {
				plan2++
			}
			plan = min(plan, plan2)
			dp[i][j] = plan
		}
	}
	return dp[len(fr)][len(sr)] <= 1
}
