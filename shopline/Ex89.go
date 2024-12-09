package shopline

/*
题目：
小红帽给外婆送糕点，由于下雨只能坐车去到外婆家。
小红帽家到外婆家之间一共有若干个车站，每个车站可能有车也可能没有车。
给定车的位置buses(假设小红帽的家在位置为0的车站，外婆家为最后一个有车的车站的位置)，开始时第一个车站的车只能移动一个车站的距离
接下来，如果上一趟车移动了b个车站的距离，那么接下来的车只能选择移动b-1，b或b+1个车站（即第一次移动了一个车站的距离，第二次可以选择移动0、1、2个车站的距离）
判断小红帽是否能顺利送糕点到外婆家。

示例1：
输入：buses = [0, 1, 3, 5, 6, 8, 12, 17]
输出：true
第一次：从位置0，坐车移动1个车站的距离，到达位置1的车站
第二次，从位置1，坐车移动2个车站的距离，到达位置3的车站
第三次，从位置3，坐车移动2个车站的距离，到达位置5的车站
第四次，从位置5，坐车移动3个车站的距离，到达位置8的车站
第五次，从位置8，坐车移动4个车站的距离，到达位置12的车站
最后，从位置12，坐车移动5个车站的距离，到达位置17的车站
即小红帽能顺利送到

示例2：
输入： buses = [0, 1, 2, 3, 4, 8, 9, 11]
输出：false
第一次：从位置0，坐车移动1个车站的距离，到达位置1的车站
第二次：从位置1，坐车移动2个车站的距离，到达位置3的车站
第三次：从位置3，可以选择移动一个车站的距离到达位置4，但是下一次最多移动2个车站，达不到位置8的车站；也可以选择移动最大距离3个车站的距离，也达不到位置8的车站
因此小红帽不能顺利送到


请用go语言实现
*/

var cache map[int]bool

func doAnswer(buses []int) bool {
	cache = make(map[int]bool, len(buses))
	for _, bus := range buses {
		cache[bus] = true
	}
	lidx := len(buses) - 1
	last := buses[lidx]
	for i := lidx - 1; i >= 0; i-- {
		v := buses[i]
		if Test(v, last-v) {
			return true
		}
	}
	return false
}

func Test(cv, step int) bool {
	if cv == 0 {
		return step == 1
	} else if cv < 0 {
		return false
	} else if step <= 0 {
		return false
	}

	// 思路：从后往前动态规划
	return (cache[cv-step] && Test(cv-step, step)) ||
		(cache[cv-(step+1)] && Test(cv-step-1, step+1)) ||
		cache[cv-(step-1)] && Test(cv-step+1, step-1)
}
