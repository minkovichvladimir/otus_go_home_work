package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
	mux      sync.Mutex
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (lc *lruCache) Set(key Key, value interface{}) bool {
	lc.mux.Lock()
	defer lc.mux.Unlock()

	item, ok := lc.items[key]
	if !ok {
		if lc.queue.Len() == lc.capacity {
			for k, listItem := range lc.items {
				if listItem == lc.queue.Back() {
					delete(lc.items, k)
					break
				}
			}

			lc.queue.Remove(lc.queue.Back())
		}

		lc.queue.PushFront(value)
		lc.items[key] = lc.queue.Front()
	} else {
		lc.queue.MoveToFront(item)
		item.Value = value
	}

	return ok
}

func (lc *lruCache) Get(key Key) (interface{}, bool) {
	lc.mux.Lock()
	defer lc.mux.Unlock()

	item, ok := lc.items[key]
	if !ok {
		return nil, false
	}
	lc.queue.MoveToFront(item)

	return item.Value, true
}

func (lc *lruCache) Clear() {
	lc.mux.Lock()
	defer lc.mux.Unlock()

	lc.items = make(map[Key]*ListItem, lc.capacity)
	lc.queue = NewList()
}
