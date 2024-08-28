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
	Key   Key
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	len   int
	first *ListItem
	// Можно и без этого элемента, но тогда при работе с последним элементом придется
	// пробегать весь список, сложность будет O(n) вместо O(1)
	last *ListItem
}

func NewList() List {
	return new(list)
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.first
}

func (l *list) Back() *ListItem {
	return l.last
}

func (l *list) PushFront(v interface{}) *ListItem {
	ni := &ListItem{
		Value: v,
		Next:  l.first,
		Prev:  nil,
	}

	if l.len == 0 {
		l.first = ni
		l.last = ni
	} else {
		l.first.Prev = ni
		l.first = ni
	}
	l.len++

	return l.first
}

func (l *list) PushBack(v interface{}) *ListItem {
	ni := &ListItem{
		Value: v,
		Next:  nil,
		Prev:  l.last,
	}

	if l.len == 0 {
		l.first = ni
		l.last = ni
	} else {
		l.last.Next = ni
		l.last = ni
	}

	l.len++

	return l.last
}

func (l *list) Remove(i *ListItem) {
	switch {
	case l.first == i:
		l.first = l.first.Next
		l.first.Prev = nil
	case l.last == i:
		l.last = l.last.Prev
		l.last.Next = nil
	default:
		i.Prev.Next = i.Next
		i.Next.Prev = i.Prev
	}

	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	switch {
	case l.first == i:
		return
	case l.last == i:
		l.last = l.last.Prev
		l.last.Next = nil
	default:
		i.Prev.Next = i.Next
		i.Next.Prev = i.Prev
	}

	i.Next = l.first
	i.Prev = nil
	l.first.Prev = i
	l.first = i
}
