func maxArea(height []int) int {
    i, j := 0, len(height)-1
    max := -1
    var current int
    for i < j {
        width := j - i
        current = height[j] * width
        if height[i] < height[j] {current = height[i] * width}
        if current > max {max = current}
        if height[i] < height[j] {
            i++
        } else {
            j--
        }  
    }
    return max
}
