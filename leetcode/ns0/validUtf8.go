package ns0

const (
	ONE   = 1 << 7  // 1000 0000
	TWO   = 3 << 6  // 1100 0000
	THREE = 7 << 5  // 1110 0000
	FOUR  = 15 << 4 // 1111 0000
	FIVE  = 31 << 3 // 1111 1000

	B00000110 = 3 << 1
	B00001110 = 7 << 1
	B00011110 = 15 << 1
	TEST      = 7 ^ 3
)

func validUtf8(data []int) bool {

	idx := 0
	n := len(data)
	byte_len := 0

	for idx < n {

		i := data[idx]
		idx++

		if byte_len != 0 {
			if i&TWO != ONE {
				return false
			}
			byte_len--
			continue
		}

		// 一字节
		if i>>7 == 0 {
			continue
		}

		tmp := i >> 3
		if tmp > B00011110 {
			return false
		}
		if tmp == B00011110 {
			byte_len = 3
			continue
		}
		tmp = i >> 4
		if tmp == B00001110 {
			byte_len = 2
			continue
		}
		tmp = i >> 5
		if tmp == B00000110 {
			byte_len = 1
			continue
		}
		return false
	}

	if byte_len != 0 {
		return false
	}
	return true
}
