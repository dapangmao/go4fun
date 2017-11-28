func judgeCircle(moves string) bool {
    var v, h = 0, 0
    for _, x := range moves {
        switch x {
            case 'U': v++
            case 'D': v--
            case 'L': h--
            default: h++
        }
    }
    return v == 0 && h == 0
}
