package medium

// 最大子数组和 https://leetcode.cn/problems/maximum-subarray/description/
// 给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
// 子数组是数组中的一个连续部分。
// 示例 1：
// 输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
// 输出：6
// 解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。
//
// 示例 2：
// 输入：nums = [1]
// 输出：1
//
// 示例 3：
// 输入：nums = [5,4,-1,7,8]
// 输出：23
func maxSubArray(nums []int) int {
	sum, max := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		cur := nums[i]
		if cur >= 0 {
			if sum < 0 {
				sum = cur
			} else {
				sum += cur
			}
		} else {
			if cur+sum > 0 {
				sum += cur
			} else {
				sum = cur
			}
		}
		if sum > max {
			max = sum
		}
	}
	return max
}

func maxSubArray2(nums []int) int {
	pre, maxVal := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		cur := nums[i]
		pre = max(cur, pre+cur)
		maxVal = max(pre, maxVal)
	}
	return maxVal
}
