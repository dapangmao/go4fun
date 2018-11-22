type interval struct {
    start int
    end int
}

func maxChunksToSorted(arr []int) int {
    var intervals []*interval
    for i, x := range arr {
        start, end := i, x
        if i > x {
            start, end = x, i
        }
        intervals = append(intervals, &interval{start, end})
    }
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i].start < intervals[j].start
    })
    res := []*interval{intervals[0]}
    for i:=1; i<len(intervals); i++ {
        current, prev := intervals[i], res[len(res)-1]
        if current.start < prev.end {
            if current.end > prev.end {prev.end = current.end}
        } else {
            res = append(res, current)
        }
    }
    return len(res)
}
