func largestSumOfAverages(A []int, K int) float64 {
    n := len(A)
    if n < K {return 0.0}
    sort.Ints(A)
    i := n-1
    var res float64
    for ; K>1; K-- {
        res += float64(A[i])
        i--
    }
    var mean float64 
    var div = float64(i + 1)
    for ; i>=0; i-- {
        mean += float64(A[i])
    }
    res += mean / div
    return res
}
