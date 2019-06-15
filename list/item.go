package list
import (
	"errors"
)
/* 
Так как я не хотел, чтобы отдельный элемент смотрел наверх, то пришлось добавить булево поле-чек на удаление
Из-за этого при вызове методов списка вначале происходит сверка не был ли первый/последний элемент удален 
*/

//item элемент списка
type item struct {
	value interface{}
	next *item
	prev *item
	delete bool
}

//newItem возвращает указатель на новый элемент
func newItem(value interface{}) *item{
	return &item{value:value}
}

//Next возвращает указатель на следующий элемент
func (i *item) Next() *item {
		return i.next
}

//Prev возвращает указатель на предыдущий элемент
func (i *item) Prev() *item {
	return i.prev
}

/* Value получение значения. 
Мне кажется здесь можно обойтись без ошибки, т.к. item скрыт и обратиться к нему можно только после 
получения элемента из методов списка Last() и Next(), которые проверяют на nil  */
func (i *item) Value() (interface{},error) {
	if i == nil {
		return nil,errors.New("Element is nil")
	}
	return i.value,nil
}

//Remove "удаляет" элемент. Меняет указатели у соседних элементов, ставит флаг удаления
func (i *item) Remove() {
	if i.prev != nil{
		i.prev.next = i.next
	}
	if i.next != nil{
		i.next.prev = i.prev
	}
	i.delete = true
}