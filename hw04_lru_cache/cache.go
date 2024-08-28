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

type KeyValue struct {
	Key   Key
	Value interface{}
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
			if kv, kvOk := lc.queue.Back().Value.(KeyValue); kvOk {
				delete(lc.items, kv.Key)
			}

			lc.queue.Remove(lc.queue.Back())
		}

		lc.queue.PushFront(KeyValue{
			Key:   key,
			Value: value,
		})
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

	if kv, kvOk := item.Value.(KeyValue); kvOk {
		return kv.Value, true
	}

	return item.Value, true
}

func (lc *lruCache) Clear() {
	lc.mux.Lock()
	defer lc.mux.Unlock()

	lc.items = make(map[Key]*ListItem, lc.capacity)
	lc.queue = NewList()
}
