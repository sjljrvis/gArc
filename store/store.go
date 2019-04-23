package store

import (
	"container/list"
	"errors"
	"sync"

	types "github.com/sjljrvis/gArch/types"
)

type storeItem struct {
	key     string
	value   *types.Peer
	element *list.Element
}

type Store struct {
	order *list.List
	items map[string]*storeItem
	limit int
	mutex *sync.Mutex
}

func Init(limit int) *Store {
	return &Store{order: list.New(), items: make(map[string]*storeItem), limit: limit, mutex: &sync.Mutex{}}
}

func (c *Store) Get(key string) (*types.Peer, error) {
	c.mutex.Lock()
	if item, exists := c.items[key]; exists {
		c.mutex.Unlock()
		return item.value, nil
	}
	return &types.Peer{Conn: nil, Msg: nil, Active: false, IP: ""}, errors.New("Key not found")
}

func (c *Store) Set(key string, value *types.Peer) (*types.Peer, error) {
	c.mutex.Lock()

	item := &storeItem{key: key, value: value}
	item.element = c.order.PushFront(item)
	c.items[key] = item

	c.mutex.Unlock()
	return item.value, nil
}

func (c *Store) All() map[string]*storeItem {
	return c.items
}
