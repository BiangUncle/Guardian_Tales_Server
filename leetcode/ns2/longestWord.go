package ns2

import "sort"

func longestWord(words []string) string {
	sort.Slice(words, func(i, j int) bool {
		if len(words[i]) < len(words[j]) {
			return true
		}
		if len(words[i]) == len(words[j]) {
			return words[i] < words[j]
		}
		return false
	})

	word_map := make(map[string]int)
	word_map[""] = 1

	ret := ""
	for _, word := range words {
		pre := word[:len(word)-1]
		_, ok := word_map[pre]
		if ok {
			if len(word) > len(ret) {
				ret = word
			}
			word_map[word] = 1
		}
	}
	return ret
}
