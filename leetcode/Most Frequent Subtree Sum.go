func findFrequentTreeSum(root *TreeNode) []int {
    var res []int
    dic := make(map[int]int)
    
    var dfs func(root *TreeNode) int
    dfs = func(root *TreeNode) int {
        if root == nil {return 0}
        left := dfs(root.Left)
        right := dfs(root.Right)
        var current = left + root.Val + right
        dic[current]++
        return current
    }
    dfs(root)
    freq := -1
    for k, v := range dic {
        if v > freq {
            res = []int{k}
            freq = v
        } else if v == freq {
            res = append(res, k)
        }
    }
    return res
}
