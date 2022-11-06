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
	sliceWithItems []ListItem
}

func (l *list) Len() int {
	return len(l.sliceWithItems)
}

func (l *list) Front() *ListItem {
	length := l.Len()
	if length < 1 {
		return nil
	}

	return &(l.sliceWithItems[0])
}

func (l *list) Back() *ListItem {
	length := l.Len()
	if length < 1 {
		return nil
	}

	return &(l.sliceWithItems[length-1])
}

func (l *list) pushFrontNewItem(newItem *ListItem) {
	length := l.Len()

	if length > 0 {
		next := &l.sliceWithItems[1]
		next.Prev = newItem
		newItem.Next = next
	}

	l.sliceWithItems = append([]ListItem{*newItem}, l.sliceWithItems...)
}

func (l *list) PushFront(v interface{}) *ListItem {
	newItem := ListItem{Value: v}
	l.pushFrontNewItem(&newItem)
	return &newItem
}

func (l *list) PushBack(v interface{}) *ListItem {
	length := l.Len()
	newItem := ListItem{Value: v}

	if length > 0 {
		prev := &l.sliceWithItems[length-1]
		prev.Next = &newItem
		newItem.Prev = prev
	}

	l.sliceWithItems = append(l.sliceWithItems, newItem)
	return &newItem
}

func (l *list) removeByIndex(index int) {
	copy(l.sliceWithItems[index:], l.sliceWithItems[index+1:])
	// l.sliceWithItems[len(l.sliceWithItems)-1] = nil
	l.sliceWithItems = l.sliceWithItems[:len(l.sliceWithItems)-1]
}

func (l *list) Remove(i *ListItem) {
	for index, item := range l.sliceWithItems {
		if item.Value == i.Value {
			l.removeByIndex(index)
			return
		}
	}
}

func (l *list) MoveToFront(i *ListItem) {
	for index, item := range l.sliceWithItems {
		if item.Value == i.Value {
			l.removeByIndex(index)
			l.pushFrontNewItem(&item)
			return
		}
	}
}

func NewList() List {
	return new(list)
}

func Rrr() {
	NewList()
}
