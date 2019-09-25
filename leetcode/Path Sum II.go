/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) [][]int {
    var res [][]int
    var dfs func(root *TreeNode, path []int, total int) 
    dfs = func(root *TreeNode, path []int, total int) {
        if root == nil {return}
        newTotal := total - root.Val
        newPath := append(append([]int{}, path... ), root.Val)

        if root.Left == nil && root.Right == nil  {
            if newTotal == 0 {res = append(res, newPath)}
            return
        }
        dfs(root.Left, newPath, newTotal)
        dfs(root.Right, newPath, newTotal)
    }
    dfs(root, []int{}, sum)
    return res
}
