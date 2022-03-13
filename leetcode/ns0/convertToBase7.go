package ns0

func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}

	neg := false
	if num < 0 {
		neg = true
		num = -num
	}
	//fmt.Println(neg)

	result := make([]byte, 0)
	cur := 5764801

	for cur > 0 {
		if num >= cur {

			bit := 0
			for num >= cur {
				bit++
				num -= cur
			}
			result = append(result, byte('0'+bit))
			cur /= 7
		} else {
			if len(result) != 0 {
				result = append(result, '0')
			}
			cur /= 7
		}
	}

	if neg {
		return "-" + string(result)
	}
	return string(result)
	/*
		7 = 1
		49 = 2
		343 = 3
		2401 = 4
		16807 = 5
		117549 = 6
		823543 = 7
		5764801 = 8
		10000000
		40353607 = 9
	*/
}
