package double_linked_list

const FirstItemInList = 1

type List struct {
	head *Item
	tail *Item
	len  int
}

func NewList() *List {
	return new(List)
}

func (l *List) Len() int {
	return l.len
}

func (l *List) First() *Item {
	return l.head
}

func (l *List) Last() *Item {
	return l.tail
}

func (l *List) PushFront(v interface{}) {
	item := &Item{value: v}
	if l.tail == nil {
		l.tail = item
	} else {
		l.head.prev = item
		item.next = l.head
	}

	l.head = item
	l.len++
}

func (l *List) PushBack(v interface{}) {
	item := &Item{value: v}
	if l.head == nil {
		l.head = item
	} else {
		l.tail.next = item
		item.prev = l.tail
	}

	l.tail = item
	l.len++
}

func (l *List) Remove(v interface{}) bool {
	counter := 1
	i := l.First()
	if i.Value() == v {
		l.remove(i, counter)

		return true
	}

	for {
		i = i.Next()
		counter++

		if i == nil {
			return false
		}

		if i.Value() == v {
			l.remove(i, counter)

			return true
		}
	}
}

func (l *List) remove(i *Item, counter int) {
	itemPrev := i.Prev()
	if itemPrev != nil {
		itemPrev.next = i.Next()
	}

	itemNext := i.Next()
	if itemNext != nil {
		itemNext.prev = i.Prev()
	}

	if counter == FirstItemInList {
		l.head = itemNext
	}

	if counter == l.len {
		l.tail = itemPrev
	}

	i.next = nil
	i.prev = nil
	i.value = nil
	l.len--
}

type Item struct {
	next  *Item
	prev  *Item
	value interface{}
}

func (i *Item) Value() interface{} {
	if i.value != nil {
		return i.value
	}

	return nil
}

func (i *Item) Next() *Item {
	if p := i.next; p != nil {
		return p
	}

	return nil
}

func (i *Item) Prev() *Item {
	if p := i.prev; p != nil {
		return i.prev
	}

	return nil
}
