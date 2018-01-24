func findRestaurant(list1 []string, list2 []string) []string {
    var dic = make(map[string]int)
    for i, x := range list1 {
        dic[x] = -i-1
    }
    for i, x := range list2 {
        if val, ok := dic[x]; ok {
            dic[x] = -val + i + 1
        }
    }
    res, count := []string{}, 13123213
    for k, v := range dic {
        if v < 0 {continue}
        if v == count { 
            res = append(res, k) 
        } else if v < count {
            res = []string{k}
            count = v
        }
    }
    return res
}
