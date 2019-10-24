/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
    res := []*TreeNode{root}
    dict := make(map[int]bool)
    for _, d := range to_delete {
        dict[d] = true
    }
    
    var dfs func(root, parent *TreeNode, isLeft bool)
    
    dfs = func(root, parent *TreeNode, isLeft bool) {
        if root == nil {return}
        left, right := root.Left, root.Right
        if dict[root.Val] {
            if isLeft {
                parent.Left = nil
            } else {
                parent.Right = nil
            }
            if left != nil && !dict[left.Val] {res = append(res, left)}
            if right != nil && !dict[right.Val] {res = append(res, right)}            
        } 
        dfs(left, root, true)
        dfs(right, root, false)
    }
    
    dfs(root, nil, true)
    return res
    
}
