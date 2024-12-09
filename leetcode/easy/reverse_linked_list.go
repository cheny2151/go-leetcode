package easy

// https://leetcode.cn/problems/reverse-linked-list/
// 给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。

type ListNode struct {
	Val  int
	Next *ListNode
}

// 递归
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	reverse := doReverse(head, head.Next)
	head.Next = nil
	return reverse
}

func doReverse(pre, cur *ListNode) *ListNode {
	if cur.Next == nil {
		cur.Next = pre
		return cur
	}
	head := doReverse(cur, cur.Next)
	cur.Next = pre
	return head
}

// 遍历
func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var pre *ListNode = nil
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
