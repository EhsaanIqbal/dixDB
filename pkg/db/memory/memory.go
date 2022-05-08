package memory

type MemoryStore struct {
	store map[string]string
}

// Initializes a memory store
func NewMemoryStore() *MemoryStore {
	ms := &MemoryStore{store: map[string]string{}}
	return ms
}

// Get -> returns corresponding value when provided a key
func (ms MemoryStore) Get(key string) string {
	return ms.store[key]
}

// Set -> stores key-value pair
func (ms MemoryStore) Set(key string, value string) {
	ms.store[key] = value
}

// Delete -> deletes key-value pair
func (ms MemoryStore) Delete(key string) {
	delete(ms.store, key)
}
