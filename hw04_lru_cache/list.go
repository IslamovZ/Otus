package hw04lrucache

type List interface {
	Show() []ListItem
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

func (l *list) Show() []ListItem {
	return l.sliceWithItems
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
	newItem.Prev = nil
	l.sliceWithItems = append([]ListItem{*newItem}, l.sliceWithItems...)
	l.setPointers(0, 1)
}

func (l *list) PushFront(v interface{}) *ListItem {
	newItem := ListItem{Value: v}
	l.pushFrontNewItem(&newItem)
	return &l.sliceWithItems[0]
}

func (l *list) PushBack(v interface{}) *ListItem {
	l.sliceWithItems = append(l.sliceWithItems, ListItem{Value: v})

	length := l.Len()
	l.setPointers(length-2, length-1)

	return &l.sliceWithItems[length-1]
}

func (l *list) removeByIndex(index int) {
	copy(l.sliceWithItems[index:], l.sliceWithItems[index+1:])
	l.sliceWithItems = l.sliceWithItems[:len(l.sliceWithItems)-1]
	l.setPointers(index-1, index)
	l.setPointers(index-2, index-1)
	l.setPointers(index, index+1)
}

func (l *list) setPointers(firstIndex int, secondIndex int) {
	length := l.Len()

	if length == 0 {
		return
	}

	var firstItem, secondItem *ListItem

	if firstIndex >= 0 && firstIndex < length {
		firstItem = &l.sliceWithItems[firstIndex]
	}
	if secondIndex >= 0 && secondIndex < length {
		secondItem = &l.sliceWithItems[secondIndex]
	}

	if firstItem != nil {
		firstItem.Next = secondItem
	}
	if secondItem != nil {
		secondItem.Prev = firstItem
	}
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
