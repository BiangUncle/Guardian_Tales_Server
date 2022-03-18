package dp

type CurNode struct {
	Dis   int
	Price int
}

//func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
//	price_dis := make([][]int, n)
//	for i := 0; i < n; i++ {
//		price_dis[i] = make([]int, n)
//	}
//
//	for _, flight := range flights {
//		price_dis[flight[0]][flight[1]] = flight[2]
//	}
//
//	cur := src
//	curnodes := make(map[int]*CurNode, 0)
//
//	next_pries := 100000
//	next_idx := -1
//
//	for i := 0; i < n; i++ {
//		p := price_dis[cur][i]
//		if p == 0 {
//			continue
//		}
//
//		node := &CurNode{Dis: i, Price: p}
//		curnodes[i] = node
//		if p < next_pries {
//			next_pries = p
//			next_idx = i
//		}
//	}
//
//	cur = next_idx
//
//	for i := 0; i < n; i++ {
//		p := price_dis[cur][i]
//		if p == 0 {
//			continue
//		}
//
//		if node, ok := curnodes[i]; ok {
//
//		}
//
//		node := &CurNode{Dis: i, Price: p}
//		curnodes[i] = node
//		if p < next_pries {
//			next_pries = p
//			next_idx = i
//		}
//	}
//}
