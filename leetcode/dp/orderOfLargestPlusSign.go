package dp

import "fmt"

func orderOfLargestPlusSign(n int, mines [][]int) int {
	up := make([][]int, n)
	left := make([][]int, n)
	sign := make([][]bool, n)
	test := make([][]int, n)

	for i := 0; i < n; i++ {
		up[i] = make([]int, n)
		left[i] = make([]int, n)
		test[i] = make([]int, n)
		sign[i] = make([]bool, n)
		for j := 0; j < n; j++ {
			sign[i][j] = true
		}
	}

	for _, point := range mines {
		sign[point[0]][point[1]] = false
	}

	ret := 0

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if !sign[i][j] {
				continue
			}

			up[i][j] = 1
			left[i][j] = 1

			if i != 0 {
				up[i][j] += up[i-1][j]
			}

			if j != 0 {
				left[i][j] += left[i][j-1]
			}

			re := HaveCross(&left, &up, i, j, n)
			test[i][j] = re
			if re > ret {
				ret = re
			}

		}
	}

	for i := 0; i < n; i++ {
		fmt.Println(sign[i])
	}
	fmt.Println("=============")
	for i := 0; i < n; i++ {
		fmt.Println(up[i])
	}
	fmt.Println("=============")
	for i := 0; i < n; i++ {
		fmt.Println(left[i])
	}
	fmt.Println("=============")
	for i := 0; i < n; i++ {
		fmt.Println(test[i])
	}

	return 1
}

func HaveCross(left, up *[][]int, i, j, n int) int {

	k := 1

	for {
		if i-k >= 0 && j+k < n {
			if (*up)[i][j] >= k+1 &&
				(*up)[i-k][j] >= k+1 &&
				(*left)[i-k][j] >= k+1 &&
				(*left)[i-k][j+k] >= k+1 {
				k++
				continue
			}
			break
		}
		break
	}

	return k
}
