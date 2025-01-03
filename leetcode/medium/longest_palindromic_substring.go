package medium

// 最长回文子串 https://leetcode.cn/problems/longest-palindromic-substring/description/
// 给你一个字符串 s，找到 s 中最长的回文子串（如果字符串向前和向后读都相同，则它满足 回文性。）
//
// 示例 1：
// 输入：s = "babad"
// 输出："bab"
// 解释："aba" 同样是符合题意的答案。
// 示例 2：
// 输入：s = "cbbd"
// 输出："bb"
func longestPalindrome(s string) string {
	var targetRunes []rune
	targetLen := 0
	chars := []rune(s)
	for i := 0; i < len(chars); i++ {
		c := chars[i]
		for j := i + targetLen; j < len(chars); j++ {
			if chars[j] == c && checkPalindromic(chars, i, j) {
				targetLen = j - i + 1
				targetRunes = make([]rune, targetLen)
				for z := 0; z < targetLen; z++ {
					targetRunes[z] = chars[i+z]
				}
			}
		}
	}
	return string(targetRunes)
}

func checkPalindromic(chars []rune, x, y int) bool {
	for x < y {
		if chars[x] != chars[y] {
			return false
		}
		x++
		y--
	}
	return true
}

// 官方题解: 中心扩展算法
func longestPalindrome2(s string) string {
	if s == "" {
		return ""
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		left1, right1 := expandAroundCenter(s, i, i)
		left2, right2 := expandAroundCenter(s, i, i+1)
		if right1-left1 > end-start {
			start, end = left1, right1
		}
		if right2-left2 > end-start {
			start, end = left2, right2
		}
	}
	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) (int, int) {
	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {
	}
	return left + 1, right - 1
}
