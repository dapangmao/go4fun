var (
	arabs  = []int{1000,900,500,400,100,90,50,40,10,9,5,4,1}
	romans = []string{"M","CM","D","CD","C","XC","L","XL","X","IX","V","IV","I"}
)

func intToRoman(num int) string {
	var res strings.Builder
	for  i:=0; i<len(arabs); i++ {
		for num >= arabs[i] {
			num -= arabs[i]
			res.Write([]byte(romans[i]))
		}
	}
	return res.String()
}


