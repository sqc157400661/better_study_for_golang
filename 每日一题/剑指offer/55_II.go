package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	return recur(root) != -1
}

func recur(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := recur(root.Left)
	if left == -1 {
		return -1
	}
	right := recur(root.Right)
	if right == -1 {
		return -1
	}
	if abs(left, right) < 2 {
		// 左 / 右子树的深度差≤1 则返回当前子树的深度
		return max(left, right) + 1
	} else {
		return -1
	}
}

func abs(x, y int) int {
	if x > y {
		return x - y
	} else {
		return y - x
	}
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
