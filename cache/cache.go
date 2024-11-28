package cache

import (
	"github.com/gorilla/websocket"
)

type Map struct {
	data map[string]*websocket.Conn
}

// New initialaze new map
func New() *Map {
	return &Map{
		data: make(map[string]*websocket.Conn, 2),
	}
}

// Set inserts new or update old value
func (m *Map) Set(key string, value *websocket.Conn) {
	m.data[key] = value
}

// Get selecte data
func (m *Map) Get(key string) (*websocket.Conn, bool) {
	v, ok := m.data[key]
	return v, ok
}

// Delete remove data by key
func (m *Map) Delete(key string) {
	delete(m.data, key)
}

// HasKey inspect key is exist
func (m *Map) HasKey(key string) bool {
	_, ok := m.data[key]
	return ok
}
