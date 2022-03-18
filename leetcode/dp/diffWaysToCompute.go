package dp

//
//func diffWaysToCompute(expression string) []int {
//	num := make([]int, 0)
//	op := make([]byte, 0)
//
//	di := ""
//	for i := 0; i < len(expression); i++ {
//		if expression[i] == '+' || expression[i] == '-' || expression[i] == '*' {
//			op = append(op, expression[i])
//			di_int, _ := strconv.Atoi(di)
//			num = append(num, di_int)
//			di = ""
//			continue
//		}
//		di = di + string(expression[i])
//	}
//	di_int, _ := strconv.Atoi(di)
//	num = append(num, di_int)
//
//	n := len(num)
//	dp := make([][]int, n)
//	for i := 0; i < n; i++ {
//		dp[i] = make([]int, n)
//		for j := 0; j < n; j++ {
//			dp[i][j] = -1
//		}
//	}
//	for idx, v := range num {
//		dp[idx][idx] = v
//	}
//
//	var dfs func(i, j int)
//	dfs = func(i, j int) int {
//
//	}
//}

func oprea(op uint8, i, j int) int {
	if op == '+' {
		return i + j
	}
	if op == '-' {
		return i - j
	}
	return i * j
}
