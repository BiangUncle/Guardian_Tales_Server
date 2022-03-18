package dp

func countVowelStrings(n int) int {

	slice := make([]int, 5)
	for i := 0; i < 5; i++ {
		slice[i] = 1
	}
	for count := 1; count < n; count++ {
		for i := 3; i >= 0; i-- {
			slice[i] = slice[i] + slice[i+1]
		}
	}
	return slice[0] + slice[1] + slice[2] + slice[3] + slice[4]
}
