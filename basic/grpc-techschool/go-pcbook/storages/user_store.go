package storages

import (
	"sync"

	"github.com/vuongle/grpc/entities"
)

type UserStore interface {
	Save(user *entities.User) error
	Find(username string) (*entities.User, error)
}

type InMemoryUserStore struct {
	mutex sync.RWMutex
	users map[string]*entities.User
}

func NewInMemoryUserStore() *InMemoryUserStore {
	return &InMemoryUserStore{
		users: make(map[string]*entities.User),
	}
}

func (store *InMemoryUserStore) Save(user *entities.User) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	if store.users[user.Username] != nil {
		return ErrAlreadyExists
	}

	store.users[user.Username] = user.Clone()
	return nil
}

func (store *InMemoryUserStore) Find(username string) (*entities.User, error) {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	user := store.users[username]
	if user == nil {
		return nil, nil
	}

	return user.Clone(), nil
}
