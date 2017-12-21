func canPlaceFlowers(flowerbed []int, n int) bool {
    var l = len(flowerbed)
    if n == 0 {return true}
    for i:=0; i<l; i++ {
        if flowerbed[i] == 1 {continue}
        if (i==0 || flowerbed[i-1] == 0 ) && (i == l-1 || flowerbed[i+1] == 0) {
            n--
            flowerbed[i] = 1
        }
        if n <= 0 {return true}
    }
    return false
}
