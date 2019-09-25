func largestTimeFromDigits(A []int) string {
    var ans = -1
    for i:= 0; i<4; i++ {
        for j:= 0; j<4; j++ {
            for k:=0; k<4; k++ {
                if i == j || i == k || k == j {continue}
                var l = 6 - i - j - k
                hours, mins := 10 * A[i] + A[j], 10 * A[k] + A[l]
                current := hours * 60 + mins
                if hours < 24 && mins < 60 && current > ans {
                    ans = current
                }
            }
        }
    }
    if ans < 0 {return ""}
    return fmt.Sprintf("%02d:%02d", ans/60, ans%60)
}
