func findMode(root *TreeNode) []int {
    max, val := -1, []int{}
    var dfs func(root *TreeNode, freq, prev int) 
    dfs = func(root *TreeNode, freq, prev int) {
        if root == nil {return}
        if (root.Left == nil && root.Right == nil) || root.Val != prev {
            if root.Val == prev {freq++}
            if freq > max {
                val = []int{prev}
                max = freq
            } else if freq == max {
                val = append(val, freq)
            }
            freq = 1
            prev = root.Val
        } else {
            freq++
        }
        dfs(root.Left, freq, prev)
        dfs(root.Right, freq, prev)
    }
    dfs(root, 0, -1)
    return val
}
