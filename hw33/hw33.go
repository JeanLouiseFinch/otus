package hw33

import (
	"sort"
	"strings"
)

type wordCount struct{
	Word string
	Count int
}

func isExist(sl []wordCount, word string) bool {
	for i := range sl {
		if sl[i].Word == word {
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
				Count: strings.Count(source,val),
				Word:val,
			})
		}
	}
	sort.Slice(result, func(i, j int) bool { return result[i].Count > result[j].Count })
	return result[:10]
}