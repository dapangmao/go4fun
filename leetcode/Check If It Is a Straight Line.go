func checkStraightLine(coordinates [][]int) bool {
    if len(coordinates) < 2 {return true}
    first, second := coordinates[0], coordinates[1]
    slope := getSlope(first, second)
    for i:=2; i<len(coordinates); i++ {
        if getSlope(coordinates[i], coordinates[i-1]) != slope {
            return false
        }    
    }
    return true
}

func getSlope(p1 []int, p2 []int) float64 {
    return float64(p2[1] - p1[1]) / float64(p2[0] - p1[0])
}
