package list

import (
	"sync"
	"errors"
	"fmt"
)
/* 
Так как я не хотел, чтобы отдельный элемент смотрел наверх, то пришлось добавить булево поле-чек на удаление
Из-за этого при вызове методов списка вначале происходит сверка не был ли первый/последний элемент удален 
*/

//List структура "список". для обновременного доступа добавлен мьютекс
type List struct{
	first *item
	last *item
	sync.Mutex
}

//NewList Создает новый список
func NewList() *List {
	return &List{}
}
//2 функции для сверки не был ли первый/последний элемент удален. Нужно после удаление последнего элемента в списке
func (l *List) checkFirst() {
	l.Lock()
	if l.first != nil && l.first.delete {
		for it := l.first; it != nil; it = l.first.next {
			if it != nil && !it.delete {
				l.first = it
				break
			}
		}
	}	
	l.Unlock()
}
func (l *List) checkLast() {
	l.Lock()
	if l.last != nil && l.last.delete {
		for it := l.last; it != nil; it = l.last.prev {
			if it != nil && !it.delete {
				l.last = it
				break
			}
		}
	}	
	l.Unlock()
}

//Print для удобного просмотра результата
func (l *List) Print() {
	l.checkFirst()
	if l.first == nil || l.first.delete {
		fmt.Println("List empty!")
		return
	}
	
	for it := l.first; it != nil; it = it.next {
		if it.prev != nil {
			fmt.Printf(" <-> (%v)",it.value)
		} else {
			fmt.Printf("\tList: (%v)",it.value)
		}
	}	
	fmt.Printf("\n")
}
//First возвращает первый элемент или ошибку если его нет
func (l *List) First() (*item,error) {
	l.checkFirst()

	if l.first == nil || l.first.delete{
		return nil,errors.New("First element nil")
	}
	return l.first,nil
}

//Last возвращает последний элемент или ошибку если его нет
func (l *List) Last() (*item,error) {
	l.checkLast()
	if l.last == nil || l.last.delete{
		return nil,errors.New("Last element nil")
	}
	return l.last,nil
}

//Len возвращает длину списка
func (l *List) Len() int{
	l.checkFirst()
	l.checkLast()
	var len int
	if l.first == nil || l.first.delete {
		return len
	}
	len++
	for it := l.first; it.next != nil; it = it.next {
		len++
	}
	return len
}

//PushFront вставка нового значения в начале списка
func (l *List) PushFront(value interface{}) {
	l.Lock()
	new := newItem(value)
	if l.first != nil && !l.first.delete{
		new.next = l.first
		l.first.prev = new
	}
	if l.last == nil || l.last.delete{
		l.last = new
	}
	l.first = new
	l.Unlock()
}

//PushBack вставка нового значения в конце списка
func (l *List) PushBack(value interface{}) {
	l.Lock()
	new := newItem(value)
	if l.last != nil && !l.last.delete{
		new.prev = l.last
		l.last.next = new
	}
	if l.first == nil || l.first.delete{
		l.first = new
	}
	l.last = new
	l.Unlock()
}