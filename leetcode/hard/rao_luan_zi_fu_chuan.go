package hard

// https://leetcode.cn/problems/scramble-string/description/
/*
使用下面描述的规则可以将字符串s1变为字符串s2，则称为和谐字符串：
1. 如果字符串的长度为1，终止
2. 如果字符串的长度 > 1 ，执行下述步骤：
-- 将字符串随机分割成两个非空的子字符串。
-- 你可以选择「交换两个子字符串」或者「保持顺序不变」
-- 两个字符串循环上述步骤，直至满足1的条件
给你两个字符串 s1 和 s2，判断s2和s1是否是和谐字符串。如果是，返回 true ；否则，返回 false 。

示例 1：
输入：s1 = "abca", s2 = "caba"
输出：true
解释：s1 上可能发生的一种情形是：
"abca" --> "ab|ca" // 分割字符串
"ab|ca" --> "ca|ab" // 交换顺序
"ca|ab" --> "c|a|a|b" // 在子字符串上递继续执行
"c|a|a|b" --> "c|a|b|a" // 第一组保持顺序不变，第二组交换顺序
算法终止，结果字符串和 s2 相同，都是 "caba", 返回 true

示例 2：
输入：s1 = "a", s2 = "a"
输出：true

示例 2：
输入：s1 = "a", s2 = "b"
输出：false
*/
func isScramble(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}
	s1len := len(s1)
	if s1len != len(s2) {
		return false
	}

	dp := make([][][]bool, s1len)

	// init dp
	for i := 0; i < s1len; i++ {
		dp[i] = make([][]bool, s1len)
		for j := 0; j < s1len; j++ {
			dp[i][j] = make([]bool, s1len)
			if s1[i] == s2[j] {
				dp[i][j][0] = true
			}
		}
	}

	for k := 1; k < s1len; k++ {
		for a := 0; a < s1len-k; a++ {
			for b := 0; b < s1len-k; b++ {
				for i := 1; i <= k; i++ {
					if dp[a][b][i-1] && dp[a+i][b+i][k-i] {
						dp[a][b][k] = true
						break
					}
					// a,b=k-(i+1)
					if dp[a][b+k-i+1][i-1] && dp[a+i][b][k-i] {
						dp[a][b][k] = true
						break
					}
				}
			}
		}
	}

	return dp[0][0][s1len-1]
}
