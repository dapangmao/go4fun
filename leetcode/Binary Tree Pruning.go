/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pruneTree(root *TreeNode) *TreeNode {
    hash := make(map[*TreeNode]bool)
    var isZeroTree func(root *TreeNode) bool
    isZeroTree = func(root *TreeNode) bool {
        if root == nil {return true}
        left, right := isZeroTree(root.Left), isZeroTree(root.Right)
        if left && right && root.Val == 0 {
            hash[root] = true
            return true
        }
        return false
    }
    isZeroTree(root)
    var dfs func(root *TreeNode) 
    dfs = func(root *TreeNode) {
        if root == nil {return}
        if hash[root.Left] {root.Left = nil}
        if hash[root.Right] {root.Right = nil}
        dfs(root.Left)
        dfs(root.Right)
    }
    dfs(root)
    return root
}
