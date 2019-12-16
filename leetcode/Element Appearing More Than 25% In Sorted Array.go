func findSpecialInteger(arr []int) int {    
    n := len(arr)
    size := n/4 
    if n % 4 != 0 {size++}
    for j:= size; j<n; j++ {
        if arr[j] == arr[j - size] {return arr[j]}
    }
    return arr[0]
}
