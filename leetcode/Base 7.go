var m = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

func convertToBase7(num int) string {
    if num == 0 {return "0"}
    var sign = ""
    if num < 0 {
        sign = "-"
        num = -num
    }
    var res []byte
    for num > 0 {
        res = append([]byte{m[num%7]}, res...)
        num /= 7
    }
    return sign + string(res)
}
