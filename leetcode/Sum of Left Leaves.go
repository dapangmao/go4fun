/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var res int

func sumOfLeftLeaves(root *TreeNode) int {
    dfs(root, false)
    return res
}

func dfs(root *TreeNode, isLeft bool) {
    if root == nil {return}
    if isLeft {res += root.Val}
    dfs(root.Left, true)
    dfs(root.Right, false)
}
