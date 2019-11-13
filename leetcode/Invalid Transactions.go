type T struct {
    name string
    time int
    money int
    city string
    raw string
}


func invalidTransactions(transactions []string) []string {
    var buf []T
    for _, t := range transactions {
        vector := strings.Split(t, ",")
        time, _ := strconv.Atoi(vector[1])
        money, _ := strconv.Atoi(vector[2])
        buf = append(buf, T{vector[0], time, money, vector[3], t})
    }
    var res []string
    for _, x := range buf {
        if x.money > 1000 {
            res = append(res, x.raw)
            goto next
        }
        for _, y := range buf {
            if x.name == y.name && x.time-y.time <=60 && y.time-x.time <=60 && x.city != y.city {
                res = append(res, x.raw)
                goto next
            } 
        }
        next:
    }
    return res
}
