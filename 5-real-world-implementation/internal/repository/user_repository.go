package repository

import (
	"errors"
	"sync"

	"github.com/danielwetan/gin-clean-architecture/internal/model"
)

type UserRepository interface {
	Save(user *model.User) error
	FindByID(id string) (*model.User, error)
	FindAll(name string) ([]*model.User, error)
}

type inMemoryUserRepository struct {
	users map[string]*model.User
	mu    sync.RWMutex
}

func NewInMemoryUserRepository() UserRepository {
	return &inMemoryUserRepository{
		users: make(map[string]*model.User),
	}
}

func (r *inMemoryUserRepository) Save(user *model.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if user.ID == "" {
		return errors.New("user ID cannot be empty")
	}

	r.users[user.ID] = user
	return nil
}

func (r *inMemoryUserRepository) FindByID(id string) (*model.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, exists := r.users[id]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (r *inMemoryUserRepository) FindAll(name string) ([]*model.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	// filter user by query param
	// users == Adam

	users := make([]*model.User, 0, len(r.users))

	if name != "" {
		for _, user := range r.users {
			if user.Name == name {
				users = append(users, user)
			}
		}
	} else {
		// No name provided, return all users
		for _, user := range r.users {
			users = append(users, user)
		}
	}

	return users, nil
}
