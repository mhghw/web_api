package store

import (
	"fmt"
	"sync"
)

func init() {
	UserDataInstance = UserData{
		UserData: make(map[string]DataRow),
	}
}

var UserDataInstance UserData

type Data struct {
	Name string
	Data interface{}
}

type DataRow struct {
	rows map[string]Data
}

type UserData struct {
	mu       sync.Mutex
	UserData map[string]DataRow
}

func (s *UserData) Add(username string, d Data) error {
	if _, has := s.UserData[username].rows[d.Name]; has {
		return fmt.Errorf("file with this name exists")
	} else {
		s.mu.Lock()
		s.UserData[username].rows[d.Name] = d
		s.mu.Unlock()
		return nil
	}
}

func (s *UserData) Get(username string, dataName string) (Data, error) {
	if d, has := s.UserData[username].rows[dataName]; !has {
		return Data{}, fmt.Errorf("file not found")
	} else {
		return d, nil
	}
}
