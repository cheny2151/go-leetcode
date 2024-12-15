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
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}
	var cur, cur2 *ListNode
	if list1.Val <= list2.Val {
		cur, cur2 = list1, &ListNode{Next: list2}
	} else {
		cur, cur2 = list2, &ListNode{Next: list1}
	}
	root := cur
	for cur.Next != nil {
		var start *ListNode
		for cur2.Next != nil && cur2.Next.Val <= cur.Next.Val {
			if start == nil {
				start = cur2.Next
			}
			cur2 = cur2.Next
		}
		if start != nil {
			cur.Next, cur = start, cur.Next
			cur2Next := cur2.Next
			cur2.Next = cur
			cur2 = &ListNode{Next: cur2Next}
		} else {
			cur = cur.Next
		}
	}

	cur.Next = cur2.Next

	return root
}
