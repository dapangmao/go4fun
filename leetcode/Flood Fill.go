func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
    queue := [][]int{[]int{sr, sc}}
    n, m := len(image), len(image[0])
    visited := make([][]bool, n)
    for i:=0; i<n; i++ {
        visited[i] = make([]bool, m)
    }
    original := image[sr][sc]
    for len(queue) > 0 {
        r, c := queue[0][0], queue[0][1]
        queue = queue[1:]
        if visited[r][c] || image[r][c] != original {continue}
        image[r][c] = newColor
        visited[r][c] = true
        if r > 0 {queue = append(queue, []int{r-1, c})}
        if r + 1 < n {queue = append(queue, []int{r+1, c})}
        if c > 0 {queue = append(queue, []int{r, c-1})}
        if c + 1 < m {queue = append(queue, []int{r, c+1})}
    }
    return image  
}
