package countword

import (
	"bufio"
	"sort"
	"strings"
)
// WordCount структура содержит строку и число ее вхождений в исходную строку
type WordCount struct{
	Word string
	Count int
}

// очистка строки от знаков препинания и приведение к нижнему регистру
func clear(s string) string {
	ex := []string{
		".",
		",",
		":",
		";",
		"?",
		"!",
		"\"",
		"(",
		")",
		"/",
	}
	for key := range ex {
		s = strings.Replace(s,ex[key]," ",-1)
	}
	return strings.ToLower(s)
}
//Top10 возвращает слайс нашей структуры. если у следующих за 10 элементом такой же счетчик вернет вернет и их тоже
func Top10(s string) []WordCount{
	s = clear(s)
	words := make(map[string]int)
	scanner := bufio.NewScanner(strings.NewReader(s))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan(){
		words[scanner.Text()]++
	}
	result := make([]WordCount,0,len(words))
	for key := range words {
		result = append(result,WordCount{
			Word: key,
			Count: words[key],
		})
	}
	sort.Slice(result,func(i,j int) bool {return result[i].Count>result[j].Count})
	if len(result)<=10 {
		return result
	}

	j:=10
	for ;j<len(result);j++{
		if result[9].Count != result[j].Count {
			break
		}
	}
	return result[:j]
}