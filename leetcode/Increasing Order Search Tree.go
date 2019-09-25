/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func NewNode(n int) *TreeNode {
    return &TreeNode{n, nil, nil}
}

func increasingBST(root *TreeNode) *TreeNode {
    if root == nil {return nil}
    buffer := dfs(root)
    prev := NewNode(buffer[0])
    res := prev
    for i:=1; i<len(buffer); i++ {
        current := NewNode(buffer[i])
        prev.Right = current
        prev = current
    }
    return res
}

func dfs(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }
    var res = append(dfs(root.Left), root.Val)
    res = append(res, dfs(root.Right)...)
    return res
}
