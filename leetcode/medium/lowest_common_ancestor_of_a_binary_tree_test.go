package medium

// 二叉树的最近公共祖先 https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/description/
// 给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。
// 百度百科中最近公共祖先的定义为：“对于有根树 T 的两个节点 p、q，最近公共祖先表示为一个节点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”
//
// 示例 1：
// 			3
// 		5		1
// 	  6	  2	   0  8
//       7 4
// 输入：root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
// 输出：3
// 解释：节点 5 和节点 1 的最近公共祖先是节点 3 。

type TreeNode3 struct {
	Val   int
	Left  *TreeNode3
	Right *TreeNode3
}

func lowestCommonAncestor(root, p, q *TreeNode3) *TreeNode3 {
	ok, node3 := find(root, p, q)
	if !ok {
		_, node3 = find(root, q, p)
	}
	return node3
}

func find(cur, p, q *TreeNode3) (bool, *TreeNode3) {
	if cur == nil {
		return false, nil
	}
	if cur.Val == p.Val {
		if find2(cur.Left, q) || find2(cur.Right, q) {
			return true, cur
		} else {
			return true, nil
		}
	}
	if ok, target := find(cur.Left, p, q); ok {
		if target != nil {
			return ok, target
		} else {
			if find2(cur.Right, q) {
				return true, cur
			} else {
				return true, nil
			}
		}
	}
	if ok, target := find(cur.Right, p, q); ok {
		if target != nil {
			return ok, target
		} else {
			if find2(cur.Left, q) {
				return true, cur
			} else {
				return true, nil
			}
		}
	}

	return false, nil
}

func find2(cur, q *TreeNode3) bool {
	if cur == nil {
		return false
	}
	if cur.Val == q.Val {
		return true
	}
	return find2(cur.Left, q) || find2(cur.Right, q)
}

// 官方题解
func lowestCommonAncestor2(root, p, q *TreeNode3) *TreeNode3 {
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return left
}
