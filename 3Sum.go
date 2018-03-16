func threeSum(num []int) [][]int {   
    sort.Ints(num)
    var res [][]int 
    for i:=0; i < len(num)-2; i++ {
        if i == 0 || (i > 0 && num[i] != num[i-1]) {
            var lo = i+1
            var hi = len(num) - 1
            var sum = 0 - num[i]
            for lo < hi {
                if (num[lo] + num[hi] == sum) {
                    res = append(res, []int{num[i], num[lo], num[hi]})
                    for lo < hi && num[lo] == num[lo+1] {lo++}
                    for lo < hi && num[hi] == num[hi-1] {hi--}
                    lo++
                    hi--
                } else if num[lo] + num[hi] < sum { 
                    lo++
                } else {
                    hi--
                }
           }
        }
    }
    return res
}
