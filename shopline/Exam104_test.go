package shopline

import (
	"math"
	"strings"
	"testing"
)

/*
题目：
你需要帮助shopline商家去优化商品的标题。商品标题由一个长度为n的字符串组成且只包含小写英文字母，好标题的定义如下：标题中每一个字符都处于连续出现至少3次的组中，
比如："aaabbb"是好标题，"aabbbb"不是好标题。
现在每次你可以执行下面2个操作之一：
选择一个下标 i（其中 0 <= i < n ）然后将该下标处的字符变为：
※该字符在字母表中 前 一个字母（前提是 title[i] != 'a' ）
※该字符在字母表中 后 一个字母（title[i] != 'z' ）

需要用最少操作次数将标题变为好标题。如果存在多个好标题，则需要返回其中字典序最小的一个。如果没有好好标题，则返回空字符串 "" 。

字典序定义：
在字符串 a 和 b 中，如果两个字符串第一个不同的字符处，字符串 a 的字母比 b 的字母在字母表里出现的顺序更早，那么我们称字符串 a 的 字典序 比 b 小 。
如果两个字符串前 min(a.length, b.length) 个字符都相同，那么较短的一个字符串字典序比另一个字符串小。

示例 1：
输入：title = "bcbc"
输出："bbbb"
解释：
"cccc" ：将 title[0] 和 title[2] 变为它们后一个字符 'c' 。
"bbbb" ：将  title[1] 和 title[3] 变为它们前一个字符 'b' 。
由于 "bbbb" 字典序比 "cccc" 小，所以返回 "bbbb" 。

示例 2：
输入：title = "aca"
输出："aaa"
解释：
无法用少于 2 个操作将字符串变为好标题。2 次操作得到好标题的方案包括：
操作 1：将 title[1] 变为 'b' ，title = "aba" 。
操作 2：将 title[1] 变为 'a' ，title = "aaa" 。
所以返回 "aaa" 。

请用go语言实现
*/

var cache2 map[int]*Slice

func Test104(t *testing.T) {
	answer104 := doAnswer104("tioffziswsvlecdapgjkmgrhfnvtaxvweqafneefxvgtxxbapamjknbsbdgc")
	if answer104 != "oooffffsssssccccjjjjmmmhhhtttwwweeefffffvvvxxxaaaakkknnncccc" {
		t.Error("must be oooffffsssssccccjjjjmmmhhhtttwwweeefffffvvvxxxaaaakkknnncccc")
	}
}

func doAnswer104(title string) string {
	cache2 = make(map[int]*Slice, 100)
	n := len(title)
	if n < 3 {
		return ""
	}

	slice := count(0, title)

	return slice.str
}

func count(idx int, title string) *Slice {
	if idx > len(title)-6 {
		return calMin(idx, len(title)-idx, title)
	}

	minStr := ""
	minDiff := math.MaxInt
	for length := 3; idx+length <= len(title)-3; length++ {
		sliceA := calMin(idx, length, title)
		if sliceA.diff >= minDiff {
			continue
		}
		var (
			sliceB *Slice
			ok     bool
		)
		if sliceB, ok = cache2[idx+length]; !ok {
			sliceB = count(idx+length, title)
			cache2[idx+length] = sliceB
		}
		diff := sliceA.diff + sliceB.diff
		str := sliceA.str + sliceB.str
		if diff < minDiff || (diff == minDiff && str < minStr) {
			minDiff = diff
			minStr = str
		}
	}

	return &Slice{diff: minDiff, str: minStr}
}

func calMin(idx, length int, title string) *Slice {
	minDiff := math.MaxInt
	minRune := uint8(0)
	for i := idx; i < idx+length; i++ {
		curDiff := 0
		for j := idx; j < idx+length; j++ {
			if j == i {
				continue
			}
			diff := int(title[j]) - int(title[i])
			if diff < 0 {
				diff = -diff
			}
			curDiff += diff
		}
		if curDiff < minDiff || (curDiff == minDiff && title[i] < minRune) {
			minDiff = curDiff
			minRune = title[i]
		}
	}
	builder := strings.Builder{}
	for i := 0; i < length; i++ {
		builder.WriteRune(rune(minRune))
	}
	return &Slice{diff: minDiff, str: builder.String()}
}

type Slice struct {
	diff int
	str  string
}
