package hub

import (
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
)

// Cache stores arbitrary data with expiration time.
type Cache struct {
	items sync.Map
	close chan struct{}
}

// An item represents arbitrary data with expiration time.
type item struct {
	data interface{}
}

// New creates a new cache
func NewCache() *Cache {
	cache := &Cache{
		close: make(chan struct{}),
	}

	return cache
}

// Get gets the value for the given key.
func (cache *Cache) Get(key interface{}) (interface{}, bool) {
	obj, exists := cache.items.Load(key)

	if !exists {
		return nil, false
	}

	item := obj.(item)

	return item.data, true
}

// Set sets a value for the given key with an expiration duration.
// If the duration is 0 or less, it will be stored forever.
func (cache *Cache) Set(key interface{}, value interface{}) {

	cache.items.Store(key, item{
		data: value,
	})
}

// Range calls f sequentially for each key and value present in the cache.
// If f returns false, range stops the iteration.
func (cache *Cache) Range(f func(key, value interface{}) bool) {
	//now := time.Now().UnixNano()

	fn := func(key, value interface{}) bool {
		item := value.(item)

		return f(key, item.data)
	}

	cache.items.Range(fn)
}

// Delete deletes the key and its value from the cache.
func (cache *Cache) Delete(key interface{}) {
	cache.items.Delete(key)
}

// Close closes the cache and frees up resources.
func (cache *Cache) Close() {
	cache.close <- struct{}{}
	cache.items = sync.Map{}
}

// =========================================

//var c = cache.New()
var c = NewCache()

func Subscribe(topic string, client *websocket.Conn) {
	clients, _ := c.Get(topic)
	if clients == nil {
		clients = make(map[*websocket.Conn]bool)
	}
	clients.(map[*websocket.Conn]bool)[client] = true

	c.Set(topic, clients)
}

func Unsubscribe(topic string, client *websocket.Conn) {

	clients, _ := c.Get(topic)
	if clients == nil {
		return
	}

	delete(clients.(map[*websocket.Conn]bool), client)
	c.Set(topic, clients)

}

func Publish(i int, topic string, data []byte) {
	clients, found := c.Get(topic)
	if found == false {
		fmt.Println("no client to send data")
		return
	}
	for c := range clients.(map[*websocket.Conn]bool) {

		if err := c.WriteMessage(i, data); err != nil {
			fmt.Println(err)
		}
	}
}
