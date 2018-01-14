func distributeCandies(candies []int) int {
    var num = len(candies) / 2
    var m = make(map[int]int)
    for _, x := range candies {
        m[x]++
    }
    if len(m) < num {return len(m)}
    return num
}
