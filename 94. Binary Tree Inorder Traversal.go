// Given the root of a binary tree, return the inorder traversal of its nodes' values.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

package main

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	tmp := []int{}
	if root.Left != nil {
		tmp = inorderTraversal(root.Left)
	}
	tmp = append(tmp, root.Val)
	if root.Right != nil {
		tmp = append(tmp, inorderTraversal(root.Right)...)
	}
	return tmp
}
