func mincostTickets(days []int, costs []int) int {
    var res = 23324324
    var dfs func(duration int, sum int, idx int)
    
    var set = make(map[[3]int]bool)
    
    dfs = func(duration int, sum int, idx int) {
        tuple := [3]int{duration, sum, idx}
        if set[tuple] == true {
            return
        } else {
            set[tuple] = true
        }
        if idx >= len(days) {
            if sum < res {res = sum}
            return
        }
        if idx > 0 && duration > 0 {
            gap := duration - days[idx] + days[idx-1]
            if gap >= 0 {
                dfs(gap, sum, idx+1)
                return
            }
        }
        dfs(0, sum+costs[0], idx+1)
        dfs(6, sum+costs[1], idx+1)
        dfs(29, sum+costs[2], idx+1)

    }
    dfs(0, 0, 0)
    return res
}
