package database

import "sync"

type DataStore struct {
	store map[string]string
	mutex sync.RWMutex
}

func NewDataStore() *DataStore {
	return &DataStore{
		store: make(map[string]string),
	}
}

func (ds *DataStore) Get(key string) (string, bool) {
	ds.mutex.RLock()
	defer ds.mutex.RUnlock()
	value, exists := ds.store[key]
	return value, exists
}

func (ds *DataStore) GetListKeys() []string {
	ds.mutex.RLock()
	defer ds.mutex.RUnlock()
	keys := make([]string, 0, len(ds.store))
	for k := range ds.store {
		keys = append(keys, k)
	}
	return keys
}

func (ds *DataStore) Put(key, value string) {
	ds.mutex.Lock()
	defer ds.mutex.Unlock()
	ds.store[key] = value
}

func (ds *DataStore) Delete(key string) bool {
	ds.mutex.Lock()
	defer ds.mutex.Unlock()
	if _, exists := ds.store[key]; exists {
		delete(ds.store, key)
		return true
	}
	return false
}
