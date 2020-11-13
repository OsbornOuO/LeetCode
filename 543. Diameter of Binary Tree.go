// Given a binary tree, you need to compute the length of the diameter of the tree. The diameter of a binary tree is the length of the longest path between any two nodes in a tree. This path may or may not pass through the root.

// Example:
// Given a binary tree
//           1
//          / \
//         2   3
//        / \
//       4   5

package main

var res = 0

func diameterOfBinaryTree(root *TreeNode) int {
	getPath(root)
	return res
}

func getPath(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := getPath(root.Left)
	right := getPath(root.Right)

	res = max(res, left+right)
	return 1 + max(left, right)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
