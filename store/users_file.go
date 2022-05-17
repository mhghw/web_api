package store

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

type User struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

var UserFileStoreInstance UserFileStore

type UserFileStore struct {
	workingDir string
}

func NewUserFileStore() {
	abs, err := filepath.Abs("./storage/users")
	if err != nil {
		log.Fatal(err)
	}

	UserFileStoreInstance.workingDir = abs

	err = os.MkdirAll(UserFileStoreInstance.workingDir, os.ModePerm)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}
}

func (s UserFileStore) AddUser(u User) error {
	if !s.UsernameAvailable(u.Username) {
		return fmt.Errorf("username already exist,please use another one")
	}
	filename := fmt.Sprintf("%v/%v.json", s.workingDir, u.Username)
	b, err := json.Marshal(u)
	if err != nil {
		return err
	}

	err = write(filename, b)

	return err
}

func (s UserFileStore) GetUser(username string) (*User, error) {
	filename := fmt.Sprintf("%v/%v.json", s.workingDir, username)

	data, err := read(filename)
	if err != nil {
		return nil, err
	}

	var u User
	err = json.Unmarshal(data, &u)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (s UserFileStore) UsernameAvailable(username string) bool {
	filename := fmt.Sprintf("%v/%v.json", s.workingDir, username)
	_, err := os.Open(filename)

	return errors.Is(err, fs.ErrNotExist)
}
