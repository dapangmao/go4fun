func isNumber(s string) bool {
    s = strings.Trim(s, " ")
    _, err := strconv.Atoi(s)
    if err == nil {return true}
    _, err = strconv.ParseFloat(s, 64)
    if err == nil {return true}
    var e = strings.Count(s, "e") 
    if strings.Count(s, ".") > 1 {return false}
    if e > 1 {return false}
    if e == 1 {
        splits := strings.Split(s, "e")
        var count int
        for i, x := range splits {
            if i == 0 {x = strings.Replace(x, ".", "", -1)}
            _, err := strconv.Atoi(x)
            if err == nil {count++}
        }
        if count == 2 {return true}
    }
    return false
}
