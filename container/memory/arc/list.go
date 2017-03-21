package arc

import (
	list2 "container/list"

	"github.com/golang-plus/caching"
)

type list struct {
	*list2.List
	Table map[string]*list2.Element
}

func (l *list) Initialize() *list {
	l.List = list2.New()
	l.Table = make(map[string]*list2.Element)

	return l
}

func (l *list) Count() int {
	return l.List.Len()
}

func (l *list) Contains(key string) bool {
	_, ok := l.Table[key]
	return ok
}

func (l *list) Get(key string) interface{} {
	if element, ok := l.Table[key]; ok {
		l.List.MoveToFront(element)
		return element.Value
	}

	return nil
}

func (l *list) Put(item *caching.Item) {
	if element, ok := l.Table[item.Key]; ok {
		element.Value = item
		l.List.MoveToFront(element)
	} else {
		element := l.List.PushFront(item)
		l.Table[item.Key] = element
	}
}

func (l *list) Discard() *caching.Item {
	element := l.List.Back()
	if element == nil {
		return nil
	}

	item := element.Value.(*caching.Item)
	l.List.Remove(element)
	delete(l.Table, item.Key)

	return item
}

func (l *list) Remove(key string) *caching.Item {
	if element, ok := l.Table[key]; ok {
		l.List.Remove(element)
		delete(l.Table, key)

		return element.Value.(*caching.Item)
	}

	return nil
}
