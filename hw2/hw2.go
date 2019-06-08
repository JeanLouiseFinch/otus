package hw2

import (
	"fmt"
	"strconv"
	"strings"
)

//symCount хранит символ и строковое значение его повторений
type symCount struct {
	sym   rune
	count string
}

//isNum проверка на 0-9
func isNum(r rune) bool {
	if r > 47 && r < 58 {
		return true
	}
	return false
}

// unpack распаковывает нашу структуру и возвращает строковое представление с учетом количества повторений символа
func (sc symCount) unpack() (string, error) {

	/*	если нужно чтобы 0 не стирал символ
		   if sc.count == "1" || sc.count == "" || sc.count == "0" {
		   		return string(sc.sym), nil
			   }
	*/
	if sc.count == "1" || sc.count == "" {
		return string(sc.sym), nil
	}
	count, err := strconv.Atoi(string(sc.count))
	if err != nil {
		return "", err
	}
	var result string
	for i := 0; i < count; i++ {
		result += string(sc.sym)
	}
	return result, nil
}

// unpackArray проходит по слайсу экземпляров нашей структуры, вызывает для каждой метод распаковки и склеивает их в результирующую строку
func unpackArray(sl []symCount) (string, error) {
	var (
		result strings.Builder
		tmp    string
		err    error
	)
	for key := range sl {
		if tmp, err = sl[key].unpack(); err != nil {
			return "", err
		}
		result.WriteString(tmp)
	}
	return result.String(), nil
}

//Unpack распаковка исходной строки
func Unpack(source string) (string, error) {
	var (
		// флаг символа экранирования
		ec bool
		// результирующий массив
		arr []symCount
		// нулевой экземпляр нашей структуры для сравнения, чтобы не записывались пустые/неправильные значения
		zero symCount
		// строка с количеством повторов. чтобы количество повторов могло быть больше 9
		countStr string
	)
	sc := symCount{}
	for _, r := range source {
		// установка флага экранирования
		if !ec && r == '\\' {
			ec = true
			continue
		}
		if ec {
			// сохранение предыдущего символа с счетчиком
			if sc != zero {
				sc.count = countStr
				arr = append(arr, sc)
			}
			// создание нового символа со счетчиком
			countStr = ""
			sc = symCount{
				sym: r,
			}
			ec = false
			continue
		}

		if isNum(r) {
			if sc == zero {
				continue
			}
			countStr += string(r)
		} else {
			// сохранение предыдущего символа с счетчиком
			if sc != zero {
				sc.count = countStr
				arr = append(arr, sc)
			}
			// создание нового символа со счетчиком
			countStr = ""
			sc = symCount{
				sym: r,
			}
		}

		ec = false
	}

	// чтобы не потерялось последнее значение
	if sc != zero {
		sc.count = countStr
		arr = append(arr, sc)
	}
	if len(arr) == 0 {
		return "", fmt.Errorf("Неправильная строка")
	}

	// вызывает функцию распаковки слайса нашей структуры
	return unpackArray(arr)
}