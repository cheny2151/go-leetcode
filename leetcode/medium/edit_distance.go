package medium

import "math"

// https://leetcode.cn/problems/edit-distance/description/
// 72. 编辑距离
// 给你两个单词 word1 和 word2， 请返回将 word1 转换成 word2 所使用的最少操作数。
//
// 你可以对一个单词进行如下三种操作：
// 插入一个字符
// 删除一个字符
// 替换一个字符
// 示例 1：
// 输入：word1 = "horse", word2 = "ros"
// 输出：3
// 解释：
// horse -> rorse (将 'h' 替换为 'r')
// rorse -> rose (删除 'r')
// rose -> ros (删除 'e')
//
// 示例 2：
// 输入：word1 = "intention", word2 = "execution"
// 输出：5
// 解释：
// intention -> inention (删除 't')
// inention -> enention (将 'i' 替换为 'e')
// enention -> exention (将 'n' 替换为 'x')
// exention -> exection (将 'n' 替换为 'c')
// exection -> execution (插入 'u')

// 递归
func minDistance(word1 string, word2 string) int {
	fr := []rune(word1)
	sr := []rune(word2)
	lenf := len(fr)
	lens := len(sr)
	// dp的x，y分别指取word1和word2的字符长度，0即为不取字符；所以要+1
	dp := make([][]*int, lenf+1)

	if lenf == 0 || lens == 0 {
		return int(math.Abs(float64(lenf) - float64(lens)))
	}

	// init
	for i := 0; i < lenf+1; i++ {
		dp[i] = make([]*int, lens+1)
		i2 := i
		dp[i][0] = &i2
	}
	for i := 1; i < lens+1; i++ {
		i2 := i
		dp[0][i] = &i2
	}

	return cal(fr, sr, dp, lenf, lens)
}

func cal(fr, sr []rune, dp [][]*int, x, y int) int {
	cur := dp[x][y]
	if cur != nil {
		return *cur
	}
	planA := cal(fr, sr, dp, x-1, y) + 1
	planB := cal(fr, sr, dp, x, y-1) + 1
	planC := cal(fr, sr, dp, x-1, y-1)
	if fr[x-1] != sr[y-1] {
		planC++
	}
	v := min(planA, min(planB, planC))
	dp[x][y] = &v
	return v
}

// 遍历
func minDistance2(word1 string, word2 string) int {
	fr := []rune(word1)
	sr := []rune(word2)
	lenf := len(fr)
	lens := len(sr)
	dp := make([][]int, lenf+1)

	if lenf == 0 || lens == 0 {
		return int(math.Abs(float64(lenf) - float64(lens)))
	}

	// init
	for i := 0; i < lenf+1; i++ {
		dp[i] = make([]int, lens+1)
		dp[i][0] = i
	}
	for i := 1; i < lens+1; i++ {
		dp[0][i] = i
	}

	for x := 1; x < lenf+1; x++ {
		for y := 1; y < lens+1; y++ {
			planA := dp[x-1][y] + 1
			planB := dp[x][y-1] + 1
			planC := dp[x-1][y-1]
			if fr[x-1] != sr[y-1] {
				planC++
			}
			dp[x][y] = min(planA, min(planB, planC))
		}
	}
	return dp[lenf][lens]
}
