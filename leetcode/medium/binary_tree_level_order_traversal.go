package medium

// 二叉树的层序遍历 https://leetcode.cn/problems/binary-tree-level-order-traversal/description/
// 给你二叉树的根节点 root ，返回其节点值的层序遍历。 （即逐层地，从左到右访问所有节点）。
// 示例 1：
// 输入：root = [3,9,20,null,null,15,7]
// 输出：[[3],[9,20],[15,7]]
// 示例 2：
// 输入：root = [1]
// 输出：[[1]]
// 示例 3：
// 输入：root = []
// 输出：[]

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	return extractNodeVal(result, 0, root)
}

func extractNodeVal(result [][]int, deep int, node *TreeNode) [][]int {
	if node == nil {
		return result
	}
	if len(result) == deep {
		result = append(result, make([]int, 0))
	}
	ints := result[deep]
	ints = append(ints, node.Val)
	result[deep] = ints
	result = extractNodeVal(result, deep+1, node.Left)
	result = extractNodeVal(result, deep+1, node.Right)
	return result
}

// 官方题解：广度优先搜索
func levelOrder2(root *TreeNode) [][]int {
	ret := [][]int{}
	if root == nil {
		return ret
	}
	q := []*TreeNode{root}
	for i := 0; len(q) > 0; i++ {
		ret = append(ret, []int{})
		p := []*TreeNode{}
		for j := 0; j < len(q); j++ {
			node := q[j]
			ret[i] = append(ret[i], node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
	}
	return ret
}
