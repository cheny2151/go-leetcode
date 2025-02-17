package shopline

/*
在一个浪漫的电影院里，有 n 对情侣相约一起来看电影。电影院的座位是一排连续排列的，总共有 2n 个座位。
工作人员提前按照一定的安排，给每对情侣发放了对应的电影票，上面标有座位号，情侣们也都按照自己票上的座位号依次入座。
每对情侣都希望能够在看电影的时候牵到对方的手，也就是他们需要坐在一起。情侣们的编号很有规律，
第一对情侣编号是 (0, 1)，第二对是 (2, 3)，依此类推，最后一对是 (2n - 2, 2n - 1)。
然而，当大家都坐好之后，却发现并不是每对情侣都挨着坐。这可让大家有些小失落，为了能让每对情侣都可以并肩坐在一起，
工作人员决定帮忙调整座位。每次调整座位，工作人员可以选择任意两位观众，让他们站起来交换座位。
现在的问题是，工作人员最少需要进行多少次座位交换，才能让每对情侣都能开心地坐在一起，手牵着手享受这场浪漫的电影呢？

示例 1:
输入: row = [0,2,1,3]
输出: 1
解释: 只需要交换row[1]和row[2]的位置即可。

示例 2:
输入: row = [3,2,0,1]
输出: 0
解释: 无需交换座位，所有的情侣都已经可以手牵手了。
请用go语言实现
*/

func doAnswer97(couples []int) int {
	cache := make(map[int]*int)
	time := 0
	for i := 0; i < len(couples); i++ {
		pre := couples[i]
		i++
		post := couples[i]
		if v, ok := cache[pre]; ok {
			time++
			pre = *v
		}
		preTg := getTarget(pre)
		if preTg == post {
			continue
		}
		if v, ok := cache[post]; ok {
			time++
			post = *v
		}
		if preTg == post {
			continue
		}
		postTg := getTarget(post)
		cache[preTg] = &post
		cache[postTg] = &pre
	}
	return time
}

func getTarget(val int) int {
	var target int
	if val%2 == 0 {
		target = val + 1
	} else {
		target = val - 1
	}
	return target
}
