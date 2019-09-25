var dic = make(map[int][]int)
var max int

func averageOfLevels(root *TreeNode) []float64 {
    var res []float64
    dfs(root, 0)
    for k:=0; k<=max; k++ {
        v := dic[k]
        res = append(res, float64(v[1]) / float64(v[0]))
	}
	return res
}

func dfs(root *TreeNode, level int) {
	if root == nil {return}
	if val, ok := dic[level]; ok {
		dic[level] = []int{val[0] + 1, val[1] + root.Val}
	} else {
		dic[level] = []int{1, root.Val}
	}
    if level > max {max = level}
	dfs(root.Left, level+1)
	dfs(root.Right, level+1)
}
