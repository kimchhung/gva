package pm

import "sync"

// MemoryStore implements Store interface with in-memory storage
type MemoryStore struct {
	roles     map[string][]string
	users     map[string]User
	resources map[string]Resource
	mu        sync.RWMutex
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		roles:     make(map[string][]string),
		users:     make(map[string]User),
		resources: make(map[string]Resource),
	}
}

func (ms *MemoryStore) GetRole(roleName string) ([]string, bool) {
	ms.mu.RLock()
	defer ms.mu.RUnlock()
	permissions, exists := ms.roles[roleName]
	return permissions, exists
}

func (ms *MemoryStore) SetRole(roleName string, permissions []string) {
	ms.mu.Lock()
	defer ms.mu.Unlock()
	ms.roles[roleName] = permissions
}

func (ms *MemoryStore) GetUser(userID string) (User, bool) {
	ms.mu.RLock()
	defer ms.mu.RUnlock()
	user, exists := ms.users[userID]
	return user, exists
}

func (ms *MemoryStore) SetUser(userID string, user User) {
	ms.mu.Lock()
	defer ms.mu.Unlock()
	ms.users[userID] = user
}

func (ms *MemoryStore) GetResource(resourceID string) (Resource, bool) {
	ms.mu.RLock()
	defer ms.mu.RUnlock()
	resource, exists := ms.resources[resourceID]
	return resource, exists
}

func (ms *MemoryStore) SetResource(resourceID string, resource Resource) {
	ms.mu.Lock()
	defer ms.mu.Unlock()
	ms.resources[resourceID] = resource
}
