package leetcode820

import (
	"bytes"
	"sort"
	"strings"
)

func minimumLengthEncoding(words []string) int {
	var buffer bytes.Buffer
	less := func(i, j int) bool {
		return len(words[i]) > len(words[j])
	}
	sort.Slice(words, less)
	for i := 0; i < len(words); i++ {
		w := words[i] + "#"
		if !strings.Contains(buffer.String(), w) {
			buffer.WriteString(w)
		}
	}
	return buffer.Len()
}

func minimumLengthEncoding2(words []string) int {
	wordsSet := make(map[string]struct{}, 0)
	for i := 0; i < len(words); i++ {
		wordsSet[words[i]] = struct{}{}
	}
	for i := 0; i < len(words); i++ {
		for j := 1; j < len(words[i]); j++ {
			delete(wordsSet, words[i][j:])
		}
	}
	var res int
	for w := range wordsSet {
		res += len(w) + 1
	}
	return res
}
