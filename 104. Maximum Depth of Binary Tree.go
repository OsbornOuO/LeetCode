// Given a binary tree, find its maximum depth.

// The maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.

// Note: A leaf is a node with no children.

// Example:

// Given binary tree [3,9,20,null,null,15,7],

package main



func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + getMax(maxDepth(root.Left), maxDepth(root.Right))
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
