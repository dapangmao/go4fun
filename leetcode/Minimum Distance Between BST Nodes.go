/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var min = 23424234
func minDiffInBST(root *TreeNode) int {
    inorder(root)
    return min
}

func inorder(root *TreeNode) []int {
    if root == nil {return []int{}} 
    var res = inorder(root.Left)
    if len(res) > 0 && root.Val - res[len(res)-1] < min {
        min = root.Val - res[len(res)-1]
    }
    res = append(res, root.Val)
    res = append(res, inorder(root.Right)...)
    return res
}
