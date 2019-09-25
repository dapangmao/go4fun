/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 */
func merge(intervals []Interval) []Interval {
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i].Start < intervals[j].Start
    })
    var res []Interval
    for i, x := range intervals {
        n := len(res) 
        if i == 0 || x.Start > res[n-1].End {
            res = append(res, x)
        } else if x.End > res[n-1].End {
            res[n-1].End = x.End
        }        
    }
    return res
}
