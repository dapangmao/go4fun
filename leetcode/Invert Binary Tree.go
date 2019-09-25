/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
    dfs(root)
    return root
}

func dfs(root *TreeNode) {
    if root == nil {return}
    root.Left, root.Right = root.Right, root.Left
    dfs(root.Left)
    dfs(root.Right)
}
