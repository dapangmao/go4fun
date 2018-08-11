/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


func sumOfLeftLeaves(root *TreeNode) int {
    var res int
    var dfs func(root *TreeNode, isLeft bool)
    dfs = func(root *TreeNode, isLeft bool) {
        if root == nil {return}
        if isLeft && root.Left == nil && root.Right == nil {
            res += root.Val
        }
        dfs(root.Left, true)
        dfs(root.Right, false)
    }
    dfs(root, false)
    return res
}
