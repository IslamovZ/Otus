package hw04lrucache

import "fmt"

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
}

func (*lruCache) Set(key Key, value interface{}) bool {
	fmt.Println(key, value)
	return true
}

func (*lruCache) Get(key Key) (interface{}, bool) {
	fmt.Println(key)

	return "asd", true
}

func (*lruCache) Clear() {
	return
}

type cacheItem struct {
	key   Key
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    nil, //NewList(),
		items:    nil, //make(map[Key]*ListItem, capacity),
	}
}
