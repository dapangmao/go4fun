func pathSum(root *TreeNode, sum int) int {
    var res int
    var dfs func(root *TreeNode, path []int) 
    dfs = func(root *TreeNode, path []int) {
        if root == nil {return}
        newPath := make([]int, len(path)+1)
        for i := range newPath  {
            if i == len(path) {
                newPath[i] = root.Val
            } else {
                newPath[i] = path[i] + root.Val
            } 
            if newPath[i] == sum {res++}
        }
        dfs(root.Left, newPath)
        dfs(root.Right, newPath)
    }
    dfs(root, []int{})
    return res
}
