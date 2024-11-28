package im

import (
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
)

// Cache stores ws.connections.
type Cache struct {
	items sync.Map
}

// An item represents ws.connections.
type item struct {
	data interface{}
}

// New creates a new cache
func NewCache() *Cache {
	return &Cache{}
}

// Get gets the value for the given key.
// key is client id
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
	cache.items = sync.Map{}
}

var c = NewCache()

// Subscribes client from a geven topic
func subscribe(topic string, client *websocket.Conn) {
	clients, _ := c.Get(topic)
	if clients == nil {
		clients = make(map[*websocket.Conn]bool)
	}
	clients.(map[*websocket.Conn]bool)[client] = true

	c.Set(topic, clients)
}

// unsubscribe delete client from topic
func unsubscribe(topic string, client *websocket.Conn) {

	clients, _ := c.Get(topic)
	if clients == nil {
		return
	}

	delete(clients.(map[*websocket.Conn]bool), client)
	c.Set(topic, clients)

}

// publish send message to all subsecribed clients
func publish(i int, topic string, data []byte) {
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
