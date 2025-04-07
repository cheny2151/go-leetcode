package shopline

import "sort"

/*
算法中心即将举行一场户外活动，物资组购买了一批物资，现在需要两个同学帮忙搬运。
很不幸，这是两个同学都不愿意比对方多搬运太多的东西
现在已知每一袋物资的重量，且每一袋物资只能由一个同学搬运，请你尽可能帮他们平均分配，并告知他们分别需要搬运的重量
输入一个数组，表示每一袋物资的重量，输出两位同学分别需要搬运的物资重量

示例1：
input：[2, 4, 1, 1]
output: [4, 4]

示例2：
input：[3, 1, 1]
output: [3, 2]

请用go语言实现
*/

func doAnswer99(weights []int) (int, int) {
	sum := 0
	for _, w := range weights {
		sum += w
	}
	target := sum / 2

	dp := make([]bool, target+1)
	dp[0] = true

	sort.Sort(sort.Reverse(sort.IntSlice(weights)))

	for _, w := range weights {
		for j := target; j >= w; j-- {
			if dp[j-w] {
				dp[j] = true
			}
		}
	}

	max := 0
	for j := target; j >= 0; j-- {
		if dp[j] {
			max = j
			break
		}
	}

	return sum - max, max
}
