package store

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var DataStoreInstance DataStore

type Data struct {
	Name string
	Data any
}

func NewData(name string, data any) *Data {
	return &Data{
		Name: name,
		Data: data,
	}
}

type DataStore struct {
	workingDir string
}

func NewDataStore() {
	abs, err := filepath.Abs("./storage/data")
	if err != nil {
		log.Fatal(err)
	}

	DataStoreInstance.workingDir = abs

	err = os.MkdirAll(UserFileStoreInstance.workingDir, os.ModePerm)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}
}

func (s DataStore) AddData(fileName, username string, d any) (string, error) {
	data := NewData(fileName, d)

	usernameIndex := fmt.Sprintf("%v/%v", s.workingDir, username)
	filename := fmt.Sprintf("%v/%v.json", usernameIndex, data.Name)
	err := os.MkdirAll(usernameIndex, os.ModePerm)
	if err != nil && !os.IsExist(err) {
		return "", err
	}

	b, err := json.Marshal(data.Data)
	if err != nil {
		return "", err
	}

	err = write(filename, b)
	if err != nil {
		return "", err
	}

	return data.Name, nil
}

func (s DataStore) GetData(username, name string) (*Data, error) {
	filename := fmt.Sprintf("%v/%v/%v.json", s.workingDir, username, name)
	b, err := read(filename)
	if err != nil {
		return nil, err
	}

	var d Data
	d.Name = name
	err = json.Unmarshal(b, &d.Data)
	if err != nil {
		return nil, err
	}

	return &d, nil
}

func (s DataStore) GetUserData(username string) ([]Data, error) {
	usernameIndex := fmt.Sprintf("%v/%v/", s.workingDir, username)

	bs, err := readAll(usernameIndex)
	if err != nil {
		return nil, err
	}

	datas := make([]Data, 0)
	for name, b := range bs {
		var d Data
		d.Name = name

		err := json.Unmarshal(b, &d.Data)
		if err != nil {
			return nil, err
		}

		datas = append(datas, d)
	}

	return datas, nil
}
