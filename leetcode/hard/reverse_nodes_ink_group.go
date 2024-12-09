package hard

// https://leetcode.cn/problems/reverse-nodes-in-k-group/
// 给你链表的头节点 head ，每 k 个节点一组进行翻转，请你返回修改后的链表。
// k 是一个正整数，它的值小于或等于链表的长度。如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
// 你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
// 示例：
// 输入：head = [1,2,3,4,5], k = 2
// 输出：[2,1,4,3,5]

type ListNode struct {
	Val  int
	Next *ListNode
}

// 低耗时，高内存消耗
func reverseKGroup(head *ListNode, k int) *ListNode {
	var newHead *ListNode
	var newCur *ListNode
	cur := head
	cache := make([]*ListNode, k)
	for c := 0; ; c++ {
		if c == k {
			c = 0
			for i := k - 1; i >= 0; i-- {
				if newHead == nil {
					newHead = cache[i]
					newCur = newHead
				} else {
					newCur.Next = cache[i]
					newCur = newCur.Next
				}
				cache[i] = nil
			}
		}
		if cur == nil {
			for _, node := range cache {
				if node != nil {
					if newHead == nil {
						newHead = node
						newCur = newHead
					} else {
						newCur.Next = node
						newCur = newCur.Next
					}
				}
			}
			newCur.Next = nil
			break
		}
		cache[c] = cur
		cur = cur.Next
	}
	return newHead
}
