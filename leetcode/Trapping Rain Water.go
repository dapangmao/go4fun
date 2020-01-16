func trap(height []int) int {
    res := 0
    n := len(height)
    left, right := make([]int, n), make([]int, n)
    var leftMax, rightMax int
    for i:=0; i<n; i++ {
        left[i] = leftMax
        if leftMax < height[i] {leftMax = height[i]}
    }
    for i:=n-1; i>=0; i-- {
        right[i] = rightMax
        if rightMax < height[i] {rightMax = height[i]}
    }
    for i:=0; i<n; i++ {        
        current := right[i] 
        if left[i] < current {current = left[i]}
        diff := current - height[i]
        if diff > 0 {res += diff}
    }
    return res
}
