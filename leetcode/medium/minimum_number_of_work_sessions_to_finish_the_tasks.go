package medium

// 完成任务的最少工作时间段:https://leetcode.cn/problems/minimum-number-of-work-sessions-to-finish-the-tasks/description/
//
// 你被安排了n个任务。任务需要花费的时间用长度为 n 的整数数组 tasks 表示，第i个任务需要花费tasks[i]小时完成。一个工作时间段中，
// 你可以至多连续工作sessionTime个小时，然后休息一会儿。
// 你需要按照如下条件完成给定任务：
// 如果你在某一个时间段开始一个任务，你需要在 同一个 时间段完成它。
// 完成一个任务后，你可以 立马 开始一个新的任务。
// 你可以按 任意顺序 完成任务。
// 给你 tasks 和 sessionTime ，请你按照上述要求，返回完成所有任务所需要的 最少 数目的 工作时间段 。
// 测试数据保证 sessionTime 大于等于 tasks[i] 中的 最大值 。
//
// 示例 1：
//
// 输入：tasks = [1,2,3], sessionTime = 3
// 输出：2
// 解释：你可以在两个工作时间段内完成所有任务。
// - 第一个工作时间段：完成第一和第二个任务，花费 1 + 2 = 3 小时。
// - 第二个工作时间段：完成第三个任务，花费 3 小时。
// 示例 2：
//
// 输入：tasks = [3,1,3,1,1], sessionTime = 8
// 输出：2
// 解释：你可以在两个工作时间段内完成所有任务。
// - 第一个工作时间段：完成除了最后一个任务以外的所有任务，花费 3 + 1 + 3 + 1 = 8 小时。
// - 第二个工作时间段，完成最后一个任务，花费 1 小时。
// 示例 3：
//
// 输入：tasks = [1,2,3,4,5], sessionTime = 15
// 输出：1
// 解释：你可以在一个工作时间段以内完成所有任务。

func minSessions(tasks []int, sessionTime int) int {
	length := len(tasks)
	// 用二进制来表示任务是否执行，假设有3个任务，那么n-1 = 111（二进制）, n= 2^3
	n := 1 << length
	dp := make([]int, n)

	// 遍历所有集合初始化dp，符合要求的集合设置为1
	for i := 1; i < n; i++ {
		total := 0
		// 查看每一个二进制位的值是否=1，=1即为执行了任务
		for cursor := 0; cursor <= i; cursor++ {
			if i&(1<<cursor) > 0 {
				total += tasks[cursor]
			}
		}
		if total <= sessionTime {
			dp[i] = 1
		}
	}

	for i := 0; i < n; i++ {
		sp := i >> 1
		// 遍历技巧j=(j-1) & i
		for j := (i - 1) & i; j > sp; j = (j - 1) & i {
			// 当前集合[i] = 当前子集+当前子集的补集，遍历所有子集取最小值
			sub := dp[j] + dp[j^i]
			if dp[i] == 0 {
				dp[i] = sub
			} else {
				dp[i] = min(dp[i], sub)
			}
		}
	}

	return dp[n-1]
}
