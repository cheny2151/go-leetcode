package shopline

/*
题目：
苹果摆成一排,每个苹果的重量不尽相同，我们用数组 apples 来表示每一个苹果的重量。
不打乱顺序的条件下分给 K 个小朋友，所以你需要将一排苹果分 K-1 次才能得到 K 份。
为了体现公平性，要求最大化最轻的那份苹果的重量；
请找出一个最佳的分割策略，使得拿到最轻的那份重量最大化，并返回这个最轻重量。


示例1：
输入： apples = [1,2,3,4,5,6,7,8,9], K = 6
输出： 6
解释： 你可以把苹果分成 [1,2,3], [4,5], [6], [7], [8], [9]。


示例2：
输入： apples = [5,6,7,8,9,1,2,3,4], K = 9
输出： 1
解释： 只有一种办法可以把苹果分成 9 块。



请用go语言实现
*/

// doAnswer92 思路：值的二分法
// 当可以分为k组，并且每一组都>=mid时，目标值在mid的右边，否则在左边
// 边界：当二分法越界(left > right)时完成遍历，返回right
func doAnswer92(apples []int, k int) int {
	minApple := apples[0]
	sumApples := 0
	for _, weight := range apples {
		if weight < minApple {
			minApple = weight
		}
		sumApples += weight
	}

	left, right := minApple, sumApples

	for left <= right {
		mid := (left + right) / 2
		if canDivide(apples, k, mid) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return right
}

// canDivide
func canDivide(apples []int, K, minWeight int) bool {
	currentSum := 0
	groupsFormed := 0
	for _, apple := range apples {
		currentSum += apple
		if currentSum >= minWeight {
			groupsFormed++
			currentSum = 0
		}
	}
	return groupsFormed >= K
}
