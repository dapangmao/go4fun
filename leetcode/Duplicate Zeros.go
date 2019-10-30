func duplicateZeros(arr []int)  {
    var queue []int
    for j := range arr {
        if len(queue) == len(arr) {
            continue
        } else if arr[j] == 0 {
            queue = append(queue, 0, 0)
        } else {
            queue = append(queue, arr[j])
        }
        arr[j] = queue[0]
        queue = queue[1:]
    }
}
