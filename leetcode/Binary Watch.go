func readBinaryWatch(num int) []string {
    var res []string
    for i:=0; i<12; i++ {
        for j:=0; j<60; j++ {
            if num == countOne(i) + countOne(j) {
                res = append(res, fmt.Sprintf("%d:%02d", i, j))
            }
        }
    }
    return res
}

func countOne(n int) int {
    var res int 
    for n > 0 {
        if n & 1 == 1 {res++}
        n >>= 1
    }
    return res
}
