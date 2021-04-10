package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	dic := map[int]int{}
}

func recur(root_idx, left_idx, right_idx int) {
	//相等就是自己
	if left_idx > right_idx {
		return
	}
}
