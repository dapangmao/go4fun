func constructRectangle(area int) []int {
    var sqrt = int(math.Sqrt(float64(area)))
    for i:=sqrt; i>1; i-- {
        var div = area / i
        if i * div == area  {return []int{div, i}}
    }
    return []int{area, 1}
}
