package medium

// 二叉树的锯齿形层序遍历 https://leetcode.cn/problems/binary-tree-zigzag-level-order-traversal/description/
// 给你二叉树的根节点root，返回其节点值的锯齿形层序遍历 。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。
//
// 示例 1：
// 输入：root = [3,9,20,null,null,15,7]
// 输出：[[3],[20,9],[15,7]]
// 示例 2：
// 输入：root = [1]
// 输出：[[1]]

type TreeNode2 struct {
	Val   int
	Left  *TreeNode2
	Right *TreeNode2
}

func zigzagLevelOrder(root *TreeNode2) [][]int {
	if root == nil {
		return [][]int{}
	}
	return next(true, []*TreeNode2{root}, [][]int{})
}

func next(left bool, nodeArr []*TreeNode2, values [][]int) [][]int {
	nextArr := make([]*TreeNode2, 0, len(nodeArr)*2)
	curValues := make([]int, 0, len(nodeArr))
	for i := len(nodeArr) - 1; i >= 0; i-- {
		curNode := nodeArr[i]
		curValues = append(curValues, curNode.Val)
		if left {
			if curNode.Left != nil {
				nextArr = append(nextArr, curNode.Left)
			}
			if curNode.Right != nil {
				nextArr = append(nextArr, curNode.Right)
			}
		} else {
			if curNode.Right != nil {
				nextArr = append(nextArr, curNode.Right)
			}
			if curNode.Left != nil {
				nextArr = append(nextArr, curNode.Left)
			}
		}
	}
	values = append(values, curValues)
	if len(nextArr) > 0 {
		values = next(!left, nextArr, values)
	}
	return values
}
