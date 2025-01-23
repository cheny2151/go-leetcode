package medium

import (
	"slices"
)

// 全排列 https://leetcode.cn/problems/permutations/description/
// 给定一个不含重复数字的数组nums，返回其所有可能的全排列。你可以按任意顺序返回答案。
//
// 示例 1：
// 输入：nums = [1,2,3]
// 输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
// 示例 2：
// 输入：nums = [0,1]
// 输出：[[0,1],[1,0]]
// 示例 3：
// 输入：nums = [1]
// 输出：[[1]]

func permute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{nums}
	}
	return fetchComp(0, nums)
}

func fetchComp(start int, nums []int) [][]int {
	var rs [][]int
	for i := start; i < len(nums); i++ {
		nums0 := cpAndMv(start, i, nums)
		if start+1 < len(nums)-1 {
			rs = append(rs, fetchComp(start+1, nums0)...)
		} else if start+1 == len(nums)-1 {
			rs = append(rs, nums0)
		}
	}
	return rs
}

func cpAndMv(i, j int, nums []int) []int {
	if i == j {
		return nums
	}
	nums0 := slices.Clone(nums)
	nums0[i], nums0[j] = nums0[j], nums0[i]
	return nums0
}
