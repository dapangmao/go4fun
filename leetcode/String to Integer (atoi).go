func myAtoi(str string) int {
        index, sign, total := 0, 1, 0
        //1. Empty string
        if len(str) == 0 {return 0}
    
        //2. Remove Spaces
        for str[index] == ' ' && index < len(str)  {
            index++;
        }
        //3. Handle signs
        if str[index] == '+' || str[index] == '-' {
            if str[index] == '-' {sign = -1}
            index ++;
        }

        //4. Convert number and avoid overflow
        for ; index < len(str); index++ {
            var digit = int(str[index] - '0')
            if digit < 0 || digit > 9  {break}
    
            //check if total will be overflow after 10 times and add digit
            if total > math.MaxInt32/10 || total == math.MaxInt32/10 && math.MaxInt32 %10 < digit {
                if sign == 1 {return math.MaxInt32}
                return math.MinInt32
            }
            total = 10 * total + digit
        }
        return total * sign
    
}
