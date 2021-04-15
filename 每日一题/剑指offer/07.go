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
//利用原理,先序遍历的第一个节点就是根。在中序遍历中通过根 区分哪些是左子树的，哪些是右子树的
//左右子树，递归
func buildTree(preorder []int, inorder []int) *TreeNode {
	//标记中序遍历
	dic := map[int]int{}
	for k,v:= range inorder {
		dic[v] = k
	}
	return  recur(0,0,len(inorder) -1,preorder,dic)
}

/**
 * @param pre_root_idx  先序遍历的索引
 * @param in_left_idx  中序遍历的索引
 * @param in_right_idx 中序遍历的索引
 */
func recur(pre_root_idx, in_left_idx, in_right_idx int,preorder []int,in_dic map[int]int) *TreeNode {
	//相等就是自己
	if in_left_idx  > in_right_idx {
		return nil
	}
	//root_idx是在先序里面的
	root := &TreeNode{Val:preorder[pre_root_idx]}

	// 有了先序的,再根据先序的，在中序中获 当前根的索引
	idx := in_dic[preorder[pre_root_idx]]

	//左子树的根节点就是 左子树的(前序遍历）第一个，就是+1,左边边界就是left，右边边界是中间区分的idx-1
	root.Left = recur(pre_root_idx + 1, in_left_idx, idx - 1,preorder,in_dic);
	//由根节点在中序遍历的idx 区分成2段,idx 就是根
	//右子树的根，就是右子树（前序遍历）的第一个,就是当前根节点 加上左子树的数量
	// pre_root_idx 当前的根  左子树的长度 = 左子树的左边-右边 (idx-1 - in_left_idx +1) 。最后+1就是右子树的根了
	root.Right = recur(pre_root_idx + (idx-1 - in_left_idx +1)  + 1, idx + 1, in_right_idx,preorder,in_dic);
	return root
}
