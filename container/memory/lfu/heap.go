package lfu

import (
	heap2 "container/heap"

	"github.com/golang-plus/caching"
)

type item struct {
	*caching.Item
	Index int // index of Item in the heap
	Count int // accessed count
}

type items []*item

func (i items) Len() int {
	return len(i)
}

func (i items) Less(x, y int) bool {
	return i[x].Count < i[y].Count
}

func (i items) Swap(x, y int) {
	i[x], i[y] = i[y], i[x]
	i[x].Index, i[y].Index = x, y
}

func (i *items) Push(x interface{}) {
	index := len(*i)
	item := x.(*item)
	item.Index = index
	*i = append(*i, item)
}

func (i *items) Pop() interface{} {
	old := *i
	index := len(old)
	item := old[index-1]
	item.Index = -1 // for safety
	*i = old[0 : index-1]

	return item
}

type heap struct {
	Items items
	Table map[string]*item
}

func (i *heap) Initialize() *heap {
	i.Items = make(items, 0)
	heap2.Init(&i.Items)
	i.Table = make(map[string]*item)

	return i
}

func (i *heap) Count() int {
	return i.Items.Len()
}

func (i *heap) Contains(key string) bool {
	_, ok := i.Table[key]
	return ok
}

func (i *heap) Get(key string) *caching.Item {
	if item, ok := i.Table[key]; ok {
		item.Count++
		heap2.Fix(&i.Items, item.Index)
		return item.Item
	}

	return nil
}

func (i *heap) Put(v *caching.Item) {
	if element, ok := i.Table[v.Key]; ok {
		element.Item = v
	} else {
		item := &item{
			Item: v,
		}
		heap2.Push(&i.Items, item)

		i.Table[v.Key] = item
	}
}

func (i *heap) Discard() *caching.Item {
	if len(i.Items) == 0 {
		return nil
	}

	item := heap2.Pop(&i.Items).(*item)
	delete(i.Table, item.Key)

	return item.Item
}

func (i *heap) Remove(key string) {
	if element, ok := i.Table[key]; ok {
		heap2.Remove(&i.Items, element.Index)
		delete(i.Table, key)
	}
}
