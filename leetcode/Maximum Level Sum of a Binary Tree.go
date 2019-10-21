/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxLevelSum(root *TreeNode) int {
    levels := []int{-1121211}
    var dfs func(root *TreeNode, level int)
    dfs = func(root *TreeNode, level int) {
        if root == nil {return}
        if len(levels) <= level {
            levels = append(levels, root.Val)
        } else {
            levels[level] += root.Val
        }
        dfs(root.Left, level+1)
        dfs(root.Right, level+1)
    }
    dfs(root, 1)
    m := max(levels)
    for i, x := range levels {
        if x == m {return i}
    }
    return 0
}


func max(arr []int) int {
	res := -123123
	for _, x := range arr {
		if x > res {res = x}
	}
	return res
}
