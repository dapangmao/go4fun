func largestTriangleArea(points [][]int) float64 {
    var res float64 = -1.0
    n := len(points)
    for i:=0; i<n; i++ {
        for j:=i+1; j<n; j++ {
            for k:=j+1; k<n; k++ {
                var current = area(points[i], points[j], points[k])
                if current > res {res = current}
            }
        }
    }
    return res
}

func area(a, b, c []int) float64 {
    var res = 0.5 * float64(a[0]*(b[1]-c[1]) + b[0]*(c[1]-a[1]) + c[0]*(a[1]-b[1]))
    if res > 0 {return res}
    return -res  
}
