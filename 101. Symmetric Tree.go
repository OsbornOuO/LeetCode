package main

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return visit(root.Left, root.Right)
}

func visit(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}

	return visit(left.Left, right.Right) && visit(left.Right, right.Left)
}
