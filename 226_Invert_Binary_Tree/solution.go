package invertbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	right := invertTree(root.Right)
	left := invertTree(root.Left)

	root.Right, root.Left = left, right

	return root
}
