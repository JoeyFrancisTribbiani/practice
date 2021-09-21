package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	queue := []*TreeNode{}
	queue = append(queue, root)
	for len(queue) != 0 {
		temp := make([]int, 0)
		len := len(queue)
		for i := 0; i < len; i++ {
			node := queue[i]
			if node != nil {
				temp = append(temp, node.Val)
			}
			if queue[i].Left != nil {
				queue = append(queue, node.Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, temp)
	}
	return res
}
