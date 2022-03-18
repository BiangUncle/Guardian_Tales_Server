package dp

import "fmt"

func generateParenthesis(n int) []string {
	ret := make([]string, 0)
	ret_map := make(map[int][]string)

	for i := 1; i <= n; i++ {
		str := ""
		left := 0
		right := 0
		var dfs func()
		dfs = func() {
			if left < i {
				left++
				str = str + "("
				dfs()
				str = str[:len(str)-1]
			}
			if right < left {
				right++
				str = str + ")"
				dfs()
				str = str[:len(str)-1]
			}
			if left == right && left == i {
				ret_map[i] = append(ret_map[i], str)
			}
		}
	}

	return ret
}

func test(i int) {
	ret_map := make(map[int][]string)

	str := ""
	left := 0
	right := 0
	var dfs func()
	dfs = func() {
		if left < i {
			left++
			str = str + "("
			dfs()
			str = str[:len(str)-1]
			left--
		}
		if right < left {
			right++
			str = str + ")"
			dfs()
			str = str[:len(str)-1]
			right--
		}
		if left == right && left == i {
			ret_map[i] = append(ret_map[i], str)
		}
	}
	dfs()
	for k, v := range ret_map {
		fmt.Println(k, v)
	}
}
