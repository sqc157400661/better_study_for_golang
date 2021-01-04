package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var Gk int
var Result int
func kthLargest(root *TreeNode, k int) int {
	Gk = k
	dfs(root)
	return Result
}

func dfs(root *TreeNode){
	if root ==nil {
		return
	}
	// 遍历右子树
	dfs(root.Right)
	if Gk ==0 {
		return
	}
	Gk --
	if Gk ==0 {
		Result = root.Val
	}
	// 遍历左子树
	dfs(root.Left)
}