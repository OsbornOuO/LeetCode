package main

func isSameTree(p *TreeNode, q *TreeNode) bool {
	return verify(p, q)
}

func verify(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}

	if verify(p.Left, q.Left) && verify(p.Right, q.Right) {
		return true
	}
	return false
}
