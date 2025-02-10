package hard

import (
	"fmt"
	"testing"
)

// 接雨水 https://leetcode.cn/problems/trapping-rain-water/description/
// 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
//
// 示例 1：
//       1
//   10001101
// 101101111111
// 输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
// 输出：6
// 解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（0表示雨水）。
// 示例 2：
// 输入：height = [4,2,0,3,2,5]
// 输出：9

func trap(height []int) int {
	sum, cur := 0, 0
	for {
		start, end := findV(cur, height)
		if end != -1 {
			s := height[start]
			e := height[end]
			m := min(s, e)
			for i := start; i <= end; i++ {
				if m > height[i] {
					sum += m - height[i]
				}
			}
			cur = end
		} else {
			break
		}
	}
	return sum
}

func findV(i int, height []int) (int, int) {
	start, end := -1, -1
	for i+1 < len(height) {
		if start == -1 && height[i] > height[i+1] {
			start = i
		} else if height[i] < height[i+1] && (end == -1 || height[i+1] > height[end]) {
			end = i + 1
			if height[i+1] >= height[start] {
				break
			}
		}
		i++
	}
	return start, end
}

func TestName(t *testing.T) {
	ints := []int{2, 8, 5, 5, 6, 1, 7, 4, 5}
	fmt.Println(trap2(ints))
}

func trap2(height []int) int {
	vArr := make([]int, 0)
	up := true
	sum, cur := 0, 0
out:
	for cur+1 < len(height) {
		if height[cur] > height[cur+1] && up {
			vArr = append(vArr, cur)
			up = false
		} else if height[cur] < height[cur+1] && len(vArr) > 0 {
			if height[cur+1] >= height[vArr[0]] {
				sum += sumV(vArr[0], cur+1, height)
				vArr = make([]int, 0)
				up = true
				continue
			}
			for j := 1; j < len(vArr); j++ {
				if height[cur+1] >= height[vArr[j]] {
					vArr[j] = cur + 1
					vArr = vArr[:j+1]
					cur++
					continue out
				}
			}
			vArr = append(vArr, cur+1)
			up = true
		}
		cur++
	}
	if len(vArr) > 1 {
		for i := 1; i < len(vArr); i++ {
			sum += sumV(vArr[i-1], vArr[i], height)
		}
	}
	return sum
}

func sumV(start, end int, height []int) int {
	sum := 0
	s := height[start]
	e := height[end]
	m := min(s, e)
	for i := start; i <= end; i++ {
		if m > height[i] {
			sum += m - height[i]
		}
	}
	return sum
}

// 官方题解:动态规划
// 创建两个长度为 n 的数组 leftMax 和 rightMax。对于 0≤i<n，leftMax[i] 表示下标 i 及其左边的位置中，height 的最大高度，rightMax[i] 表示下标 i 及其右边的位置中，height 的最大高度。
// 在得到数组 leftMax 和 rightMax 的每个元素值之后，对于 0≤i<n，下标 i 处能接的雨水量等于 min(leftMax[i],rightMax[i])−height[i]。遍历每个下标位置即可得到能接的雨水总量。
func trap3(height []int) (ans int) {
	n := len(height)
	if n == 0 {
		return
	}

	leftMax := make([]int, n)
	leftMax[0] = height[0]
	for i := 1; i < n; i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}

	rightMax := make([]int, n)
	rightMax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}

	for i, h := range height {
		ans += min(leftMax[i], rightMax[i]) - h
	}
	return
}

// 官方题解:双指针
func trap4(height []int) (ans int) {
	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	for left < right {
		leftMax = max(leftMax, height[left])
		rightMax = max(rightMax, height[right])
		if height[left] < height[right] {
			ans += leftMax - height[left]
			left++
		} else {
			ans += rightMax - height[right]
			right--
		}
	}
	return
}
