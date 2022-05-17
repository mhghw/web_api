package store

import "io/ioutil"

func write(filename string, data []byte) error {
	err := ioutil.WriteFile(filename, data, 0644)
	return err
}
