func escapeGhosts(ghosts [][]int, target []int) bool {
    tx, ty := target[0], target[1]
    var escape = abs(tx) + abs(ty)
    for _, g := range ghosts {
        gx, gy := g[0], g[1]
        dist := abs(gx - tx) + abs(gy - ty)
        if dist < escape {return false}
    }
    return true
}

func abs(a int) int{
    if a < 0 {return -a}
    return a
}
