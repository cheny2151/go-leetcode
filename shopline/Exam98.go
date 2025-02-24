package shopline

/*
*
N个灯泡拍成一行，编号从1到N。最初，所有的灯泡都关闭。每天只打开一个灯泡，直到N天后所有的灯泡都打开。
给你一个长度为N的灯泡数组blubs,其中blubs[i]=x表示在第i+1天，我们会把位置在x的灯泡打开，其中i从0开始，x从1开始。
给你一个整数k，请你输出在第几天恰好有两个打开的灯泡，使得他们中间正好有k个灯泡且这些灯泡全是关闭的。
如果不存在这种情况，返回-1.如果有多天出现这种情况，返回最小的天数。

样例 1:
输入:
blubs: [1,3,2]
k: 1
输出: 2
解释: 在第二天，第一和第三灯都开了,中间刚好k=1个灯泡是关闭的。

样例 2:
输入:
blubs: [1,2,3]
k: 1
输出: -1

注释 :
给定的数组范围是 [1, 20000]。
*/
func doAnswer98(blubs []int, k int) int {
	length := len(blubs)
	bitMap := NewBitMap(length)
	for i, blub := range blubs {
		bitMap.setAt(blub - 1)
		if blub+k < length && bitMap.checkAt(blub+k) {
			pass := true
			for j := blub; j < blub+k; j++ {
				if bitMap.checkAt(j) {
					pass = false
					break
				}
			}
			if pass {
				return i + 1
			}
		}
		if blub-2-k >= 0 && bitMap.checkAt(blub-2-k) {
			pass := true
			for j := blub - 2; j > blub-2-k; j-- {
				if bitMap.checkAt(j) {
					pass = false
					break
				}
			}
			if pass {
				return i + 1
			}
		}
	}
	return -1
}

type BitMap struct {
	offset []int64
}

func NewBitMap(cap int) *BitMap {
	size := cap >> 6
	if cap&63 > 0 {
		size++
	}
	return &BitMap{offset: make([]int64, size)}
}

func (bm *BitMap) setAt(index int) {
	offsetIdx := index >> 6
	bm.offset[offsetIdx] |= 1 << (index & 63)
}

func (bm *BitMap) checkAt(index int) bool {
	offsetIdx := index >> 6
	return bm.offset[offsetIdx]&(1<<(index&63)) > 0
}
