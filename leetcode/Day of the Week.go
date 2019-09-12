func dayOfTheWeek(day int, month int, year int) string {
    current := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
    return current.Weekday().String()
}
