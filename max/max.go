package max


import (
	"reflect"
)

//GetMax Возвращает максимум в слайсе. Если sl не слайс или он пустой - паника
func GetMax(sl interface{}, compare func(i, maxIdx int) bool) interface{} {

	var max interface{}

	val := reflect.ValueOf(sl)
	if val.Kind() != reflect.Slice {
		panic("Not slice!")
	}
	if val.Len() == 0 {
		panic("Len slice 0!")
	}
	
	// Принимаем что первое значение макисмум
	idx := 0
	max = val.Index(idx).Interface()

	for i := 1; i < val.Len(); i++ {
		if compare(i, idx) {
			max = val.Index(i).Interface()
			idx = i
		}
	}
	return max
}
