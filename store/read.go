package store

import (
	"os"
)

func read(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	//get file info
	info, err := file.Stat()
	if err != nil {
		return nil, err
	}

	buffer := make([]byte, info.Size())
	_, err = file.Read(buffer)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}
