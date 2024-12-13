package medium

// https://leetcode.cn/problems/merge-two-sorted-lists/description/
// 将两个升序链表合并为一个新的升序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
// 示例
// 输入：l1 = [1,2,4], l2 = [1,3,4]
// 输出：[1,1,2,3,4,4]

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	cur := list1
	for {
		var pre *ListNode
		for list2.Val <= cur.Val {
			if pre == nil {
				pre = list2
			}
			list2 = list2.Next
		}
		if pre != nil {
			cur.Next, cur,pre.Next = pre, cur.Next,
		}
	}

	return list1
}
