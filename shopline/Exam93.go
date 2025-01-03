package shopline

/*
题目：
算法中心计划开展一次团建活动，此次团建涉及到 n 个项目的预算安排，这些项目对应的金额构成了非负整数数组 budget
每个项目都存在两种可能：成本或者收益。即给每个项目的金额可以为正值或者负值。
作为团建组织者，你肩负着控制团建预算的重任，你的目标是通过合理安排各个项目的预算，统计有多少种不同的预算分配方式使得最终的总预算恰好等于核定的目标整数值target。

示例1：
输入：budget = [1, 999], target = 998
输出：1

解释：
只有1种分配方式，项目0为收益，付负值，项目1为成本，付正值，总预算为-1+999=998


示例2：
输入：budget = [1, 1, 1, 1, 1], target = 3
输出：5

解释：
（1） 项目0为收益，项目1为成本，项目2为成本，项目3为成本，项目4为成本，总预算为-1+1+1+1+1 = 3
（2） 项目0为成本，项目1为收益，项目2为成本，项目3为成本，项目4为成本，总预算为 1-1+1+1+1 = 3
（3） 项目0为成本，项目1为成本，项目2为收益，项目3为成本，项目4为成本，总预算为 1+1-1+1+1 = 3
（4） 项目0为成本，项目1为成本，项目2为成本，项目3为收益，项目4为成本，总预算为 1+1+1-1+1 = 3
（5） 项目0为成本，项目1为成本，项目2为成本，项目3为成本，项目4为收益，总预算为 1+1+1+1-1 = 3
一共有 5 种不同的预算分配方式

*/

// 核心思路：sum(budget) - target得到的数就是数组元素需要取负的总和，由于+变-所以元素要*2
// 所以问题转化为：遍历数组有多少子集（sumTarget） 元素*-1*2 = sum(budget) - target
func doAnswerEx93(budget []int, target int) int {
	sum := 0
	for _, v := range budget {
		sum += v
	}

	sub := sum - target
	count := sumTarget(0, 0, sub, budget)
	// sub为0时，什么都不取本身就是一个case
	if sub == 0 {
		count++
	}

	return count
}

func sumTarget(curSum, idx, target int, arr []int) int {
	if idx == len(arr) {
		return 0
	}

	val := arr[idx] * 2
	nextSum := curSum + val
	count := 0
	if nextSum == target {
		// 由于后续元素可能为0，所以得到目标值后还需要继续遍历
		count += 1 + sumTarget(nextSum, idx+1, target, arr)
	} else if nextSum < target {
		count += sumTarget(nextSum, idx+1, target, arr)
	}
	count += sumTarget(curSum, idx+1, target, arr)
	return count
}
