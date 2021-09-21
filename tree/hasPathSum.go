package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	sum := 0
	return dfsPath(root, sum, targetSum)
}
func dfsPath(root *TreeNode, sum int, targetSum int) bool {
	if root == nil {
		if sum == targetSum {

			return true
		} else {
			return false
		}
	}
	sum = sum + root.Val
	if root.Left == nil {
		return dfsPath(root.Right, sum, targetSum)
	}
	if root.Right == nil {
		return dfsPath(root.Left, sum, targetSum)
	}
	return dfsPath(root.Left, sum, targetSum) || dfsPath(root.Right, sum, targetSum)
}
