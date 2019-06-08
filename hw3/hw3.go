package hw3

import (
	"sort"
	"strings"
)

type wordCount struct{
	word string
	count int
}

func isExist(sl []wordCount, word string) bool {
	for i := range sl {
		if sl[i].word == word {
			return true
		}
	}
	return false
}

func Top10(source string) ([]wordCount) {
	var (
		str []string
		result []wordCount
	)
	str = strings.Split(source," ")
	result = make([]wordCount,0)
	for _,val := range str {
		if !isExist(result,val) {
			result = append(result,wordCount{
				count: strings.Count(source,val),
				word:val,
			})
		}
	}
	sort.Slice(result, func(i, j int) bool { return result[i].count > result[j].count })
	return result[:10]
}