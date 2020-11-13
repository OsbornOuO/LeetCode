// Given a binary tree, you need to compute the length of the diameter of the tree. The diameter of a binary tree is the length of the longest path between any two nodes in a tree. This path may or may not pass through the root.

// Example:
// Given a binary tree
//           1
//          / \
//         2   3
//        / \
//       4   5

package main

import "testing"

func Test_diameterOfBinaryTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				root: &TreeNode{
					Val: 1,
					// Right: &TreeNode{
					// 	Val: 2,
					// 	Right: &TreeNode{
					// 		Val: 3,
					// 	},
					// },
					// Left: &TreeNode{
					// 	Val: 4,
					// 	Right: &TreeNode{
					// 		Val: 5,
					// 		Right: &TreeNode{
					// 			Val: 6,
					// 		},
					// 	},
					// },
				},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := diameterOfBinaryTree(tt.args.root); got != tt.want {
				t.Errorf("diameterOfBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
