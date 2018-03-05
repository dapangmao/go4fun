func candy(ratings []int) int {
    var n = len(ratings)
    if n == 0 {return 0}
    if n == 1 {return 1}
    left, right := make([]int, n), make([]int, n)
    for i:=0; i<n; i++ {
        left[i] = 1
        right[i] = 1
    }
    for i:=1; i<n; i++ {
        if ratings[i] > ratings[i-1] {
            left[i] = left[i-1] + 1
        }
    }
    for i:=n-2; i>=0; i-- {
        if ratings[i] > ratings[i+1] {
            right[i] = right[i+1] + 1
        } 
    }
    var res int
    for i:=0; i<n; i++ {
        var current = left[i]
        if left[i] < right[i] {current = right[i]}
        res += current
    }
    return res
}
