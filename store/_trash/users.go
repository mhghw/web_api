package store

import (
	"fmt"
	"sync"
)

func init() {
	UserStoreInstance = UserStore{
		users: make(map[string]User),
	}
}

var UserStoreInstance UserStore

type User struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

type UserStore struct {
	mu    sync.Mutex
	users map[string]User
}

func (s *UserStore) AddUser(u User) error {
	if s.UsernameAvailable(u.Username) {
		return fmt.Errorf("username: %v already exists", u.Username)
	} else {
		s.mu.Lock()
		s.users[u.Username] = u
		s.mu.Unlock()
		return nil
	}
}

func (s *UserStore) UsernameAvailable(username string) bool {
	if _, has := s.users[username]; has {
		return true
	} else {
		return false
	}
}

func (s *UserStore) GetUser(username string) (User, error) {
	user, has := s.users[username]
	if !has {
		return User{}, fmt.Errorf("user does not exist")
	}
	return user, nil
}
