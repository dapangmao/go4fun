func missingNumber(arr []int) int {
    n := len(arr)
    diff := (arr[n-1] - arr[0]) / n
    for i:=1; i<n; i++ {
        should := arr[i-1] + diff
        if arr[i] != should {return should} 
    }
    return arr[0]
}
