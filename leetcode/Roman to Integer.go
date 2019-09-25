func romanToInt(s string) int {
    var res = 0
    if strings.Contains(s, "IV")  {res -= 2}
    if strings.Contains(s, "IX")  {res -= 2}
    if strings.Contains(s, "XL") {res -= 20}
    if strings.Contains(s, "XC") {res -= 20}
    if strings.Contains(s, "CD") {res -= 200}
    if strings.Contains(s, "CM") {res -= 200}
    for _, c := range s {
        switch c {
            case 'M': 
            res += 1000
            case 'D':
            res += 500
            case 'C':
            res += 100
            case 'L':
            res += 50
            case 'X':
            res += 10
            case 'V':
            res += 5
            case 'I':
            res += 1
        }
    }
    return res
}
