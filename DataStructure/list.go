package main

import "fmt"

func main() {
	list := NewList()
	list.add(1).add(2).add(3).add(4).add(5)
	for k, v := range list.array {
		fmt.Println(k, v)
	}
}

type List struct {
	cap         int
	len         int
	array       []int
	extendRadio int
}

func NewList() *List {
	return &List{
		cap:         4,
		len:         0,
		array:       make([]int, 0, 4),
		extendRadio: 2,
	}
}

func (l *List) add(n int) *List {
	if l.len >= l.cap {
		l.array = l.expansion()
	}
	l.array = append(l.array, n)
	l.len += 1
	return l
}
func (l *List) expansion() []int {
	l.cap *= l.extendRadio
	newArray := make([]int, 0, l.cap)
	for _, v := range l.array {
		newArray = append(newArray, v)
	}
	return newArray
}
