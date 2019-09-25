func topKFrequent(nums []int, k int) []int {
    dic := make(map[int]int)
    max := -1
    for _, num := range nums {
        dic[num]++
        if dic[num] > max {max = dic[num]}
    }
    var bucket = make([][]int, max+1)
    for k, v := range dic {
        bucket[v] = append(bucket[v], k)
    }
    var res []int
    for i:=max; i>=0; i-- {
        if len(bucket[i]) == 0 {continue}
        res = append(res, bucket[i]...)
        if len(res) == k {return res}        
    }
    return res
}
