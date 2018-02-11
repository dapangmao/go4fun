func largestNumber(nums []int) string {
    var strs []string
    for _, num := range nums {
        strs = append(strs, strconv.Itoa(num))
    }
    sort.Slice(strs, func(i, j int) bool {
        return strs[i] + strs[j] > strs[j] + strs[i]
    })
    if strs[0] == "0" {return "0"}
    return strings.Join(strs, "")
}
