func maxNumberOfBalloons(text string) int {
    res := 8888888
    dict := make(map[rune]int)
    for _, r := range text {
        for _, t := range "balon" {
            if r == t {
                dict[r]++
            }
        }
    }
    if len(dict) != 5 {return 0}
    for k, v := range dict {
        if k == rune('l') || k == rune('o') {
            if v < 2 {return 0}
            dict[k] /= 2
        }
    }
    for _, v := range dict {
        if v < res {res = v}
    }
    if res == 8888888 {return 0}
    return res
}
