func carPooling(trips [][]int, capacity int) bool {
    mileages := make([]int, 1001)
    for _, trip := range trips {
        num, start, end := trip[0], trip[1], trip[2]
        for i:=start; i<end; i++ {
            mileages[i] += num
            if mileages[i] > capacity {return false}
        }
    }
    return true
}
