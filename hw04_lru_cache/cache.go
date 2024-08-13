package hw04lrucache

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

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

type Pair struct {
	key   Key
	value interface{}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	element, keyExists := c.items[key]
	if keyExists {
		c.queue.MoveToFront(element)
		element.Value.(*Pair).value = value
	} else {
		if c.queue.Len() == c.capacity {
			c.queue.Remove(c.queue.Back())
			delete(c.items, c.queue.Back().Value.(*Pair).key)
		}
		pair := &Pair{key, value}
		element = c.queue.PushFront(pair)
		c.items[key] = element
	}
	return keyExists
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	element, keyExists := c.items[key]
	if keyExists {
		c.queue.MoveToFront(element)
	} else {
		return nil, keyExists
	}
	return element.Value.(*Pair).value, keyExists
}

func (c *lruCache) Clear() {
	c.capacity = 0
	c.queue = new(list)
	c.items = make(map[Key]*ListItem, c.capacity)
}
