package medium

// https://leetcode.cn/problems/longest-substring-without-repeating-characters/description/
// 无重复字符的最长子串
// 给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
// 示例 1:
//
// 输入: s = "abcabcbb"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
// 示例 2:
//
// 输入: s = "bbbbb"
// 输出: 1
// 解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
// 示例 3:
//
// 输入: s = "pwwkew"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//      请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

func lengthOfLongestSubstring(s string) int {
	chars := []rune(s)
	if len(chars) == 1 {
		return 1
	}
	v := 0
	preIdx := 0
	cache := make(map[rune]*int)
	for i := range chars {
		ch := chars[i]
		cv := cache[ch]
		if cv != nil && *cv >= preIdx {
			v = max(v, i-preIdx)
			preIdx = *cv + 1
		}
		i2 := i
		cache[ch] = &i2
	}
	v = max(v, len(chars)-preIdx)
	return v
}
