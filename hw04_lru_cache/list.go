package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	front *ListItem
	back  *ListItem
	len   int
}

func NewList() List {
	return new(list)
}

func (l list) Len() int {
	return l.len
}

func (l list) Front() *ListItem {
	if l.len == 0 {
		return nil
	}
	return l.front
}

func (l list) Back() *ListItem {
	if l.len == 0 {
		return nil
	}
	return l.back
}

func (l *list) PushFront(v interface{}) *ListItem {
	e := &ListItem{v, nil, nil}
	if l.front == nil {
		l.front = e
		l.back = e
	} else {
		l.front.Prev = e
		e.Next = l.front
		e.Prev = nil
		l.front = e
	}
	l.len++
	return e
}

func (l *list) PushBack(v interface{}) *ListItem {
	e := &ListItem{v, nil, nil}
	if l.back == nil {
		l.front = e
		l.back = e
		l.len++
	} else {
		l.back.Next = e
		e.Next = nil
		e.Prev = l.back
		l.back = e
	}
	l.len++
	return e
}

func (l *list) Remove(i *ListItem) {
	i.Next.Prev = i.Prev
	i.Prev.Next = i.Next
	i.Next = nil
	i.Prev = nil
	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	if i == l.Front() {
		return
	}
	if i != l.Back() {
		i.Next.Prev = i.Prev
	}
	i.Prev.Next = i.Next

	i.Next = l.front
	i.Prev = l.front.Prev

	i.Next.Prev = i
	l.front = i
}
