/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var res int

func sumNumbers(root *TreeNode) int {
    dfs(root, 0)
    return res
}

func dfs(root *TreeNode, path int) {
    if root == nil {
        return
    }
    var current = path*10 + root.Val
    if root.Left == nil && root.Right == nil {
        res += current
        return
    }
    if root.Left != nil {dfs(root.Left, current)}
    if root.Right != nil {dfs(root.Right, current)}
}
