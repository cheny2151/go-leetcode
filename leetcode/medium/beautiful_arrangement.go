package medium

import (
	"math/bits"
)

// 优美的排列 https://leetcode.cn/problems/beautiful-arrangement/description/
// 假设有从 1 到 n 的 n 个整数。用这些整数构造一个数组 perm（下标从 1 开始），只要满足下述条件之一 ，该数组就是一个优美的排列：
// perm[i] 能够被 i 整除
// i 能够被 perm[i] 整除
// 给你一个整数 n ，返回可以构造的优美排列的数量。
//
// 示例 1：
// 输入：n = 2
// 输出：2
// 解释：
// 第 1 个优美的排列是 [1,2]：
//    - perm[1] = 1 能被 i = 1 整除
//    - perm[2] = 2 能被 i = 2 整除
// 第 2 个优美的排列是 [2,1]:
//    - perm[1] = 2 能被 i = 1 整除
//    - i = 2 能被 perm[2] = 1 整除
// 示例 2：
// 输入：n = 1
// 输出：1

func countArrangement(n int) int {
	// 用n个位数的二进制表示对应index是否被占用（1代表被占用，也代表了可以被整除）
	// 然后用1<<n个元素的数组来存储结果
	// 例如111就代表了N=3的最终结果，并且111对应的数组位置values[1<<3-1]即为最终答案
	values := make([]int, 1<<n)
	values[0] = 1
	for i := 1; i < 1<<n; i++ {
		// 1的个数即为当前数字
		num := bits.OnesCount(uint(i))
		for j := 1; i <= n; i++ {
			// 判断当前位置j在当前i的二进制下是否占用(为1)并且当前num是否满足条件（放在j位置）
			if i>>(j-1)&1 == 1 && (num%j == 0 || j%num == 0) {
				values[i] += values[1<<(j-1)^i]
			}
		}
	}
	return values[1<<n-1]
}

// 官方题解
func countArrangement2(n int) int {
	f := make([]int, 1<<n)
	f[0] = 1
	for mask := 1; mask < 1<<n; mask++ {
		num := bits.OnesCount(uint(mask))
		for i := 0; i < n; i++ {
			if mask>>i&1 > 0 && (num%(i+1) == 0 || (i+1)%num == 0) {
				f[mask] += f[mask^1<<i]
			}
		}
	}
	return f[1<<n-1]
}
