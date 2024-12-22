package shopline

/*
题目：

你在星际航行过程中，到了一个遥远的星球，在这里有一种自动化的超级洗衣机。这些洗衣机不仅可以自己洗衣服，还能够通过高科技的传送装置，将衣服从一台洗衣机传送给相邻的洗衣机。
有一天，星球上的清洁工发现一个问题，就是一排 n 台超级洗衣机，每台洗衣机里装着不同数量的衣服。为了让所有洗衣机能够高效工作，清洁总管决定，必须让每台洗衣机最终装有相同数量的衣服。
于是，清洁工们找到来处远方你，来帮助解决这个问题。你观察了一下洗衣机，发现了几条规则：

.

	在每一步操作中，每台洗衣机都可以将一件衣服（最多一件）传送给它的相邻洗衣机。
	每台洗衣机同时只能与左右相邻的洗衣机传递衣服，不能直接跨距离传送。且洗衣机是直线排列非环形排列，即收尾两个洗衣机无法联通。

你的任务是帮清洁工们设计一个最优方案，计算出至少需要多少步操作，才能让所有洗衣机中的衣物数量相等。如果发现无论如何也做不到这一点，你需要及时告诉他们，这个任务根本不可能完成。

示例 1：

	输入：machines = [1,0,5]
	输出：3
	解释：
	第一步:    1     0 <-- 5    =>    1     1     4
	第二步:    1 <-- 1 <-- 4    =>    2     1     3
	第三步:    2     1 <-- 3    =>    2     2     2

示例 2：

	输入：machines = [0,3,0]
	输出：2
	解释：
	第一步:    0 <-- 3     0    =>    1     2     0
	第二步:    1     2 --> 0    =>    1     1     1

示例 3：

	输入：machines = [0,2,0]
	输出：-1
	解释：
	不可能让所有三个洗衣机同时剩下相同数量的衣物。
*/
func doAnswer9101(washer []int) int {
	length := len(washer)
	sum := 0
	for i := 0; i < length; i++ {
		sum += washer[i]
	}
	avg := sum / length
	if avg*length != sum {
		return -1
	}

	start := 0
	count := 0
	for {
		pre := start * avg
		change := -1
		for i := start; i < length; i++ {
			pre += washer[i]
			expire := (i + 1) * avg
			if pre == expire {
				if change == -1 {
					start = i + 1
				} else {
					if change == 1 {
						washer[i]++
					} else if change == 2 {
						washer[i]--
					}
					change = 0
				}
				continue
			} else if pre > expire {
				if change != 1 {
					washer[i]--
				}
				if change == 2 {
					change = 0
				} else {
					change = 1
				}
			} else if pre < expire {
				if change == 1 {
					washer[i]++
					change = 0
				}
				if change <= 0 {
					washer[i]++
					change = 2
				}
			}
		}
		if change >= 0 {
			count++
		} else {
			break
		}
	}
	return count
}
