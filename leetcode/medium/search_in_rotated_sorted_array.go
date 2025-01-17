package medium

// 搜索旋转排序数组 https://leetcode.cn/problems/search-in-rotated-sorted-array/description/
// 整数数组 nums 按升序排列，数组中的值互不相同 。
// 在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转，使数组变为
// [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。
// 例如， [0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为 [4,5,6,7,0,1,2] 。
// 给你 旋转后的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回 -1 。
// 你必须设计一个时间复杂度为 O(log n) 的算法解决此问题。
//
// 示例 1：
// 输入：nums = [4,5,6,7,0,1,2], target = 0
// 输出：4
//
// 示例 2：
// 输入：nums = [4,5,6,7,0,1,2], target = 3
// 输出：-1
//
// 示例 3：
// 输入：nums = [1], target = 0
// 输出：-1
func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	return searchPart(nums, left, right, target)
}

func searchPart(nums []int, left, right, target int) int {
	if left > right {
		return -1
	}
	lfv := nums[left]
	if target == lfv {
		return left
	}
	rgv := nums[right]
	if target == rgv {
		return right
	}
	mid := (left + right) / 2
	midv := nums[mid]
	if target == midv {
		return mid
	}
	if midv > lfv {
		if target >= lfv && target < midv {
			return searchPart(nums, left, mid-1, target)
		}
		return searchPart(nums, mid+1, right, target)
	} else {
		if target > midv && target <= rgv {
			return searchPart(nums, mid+1, right, target)
		}
		return searchPart(nums, left, mid-1, target)
	}
}
