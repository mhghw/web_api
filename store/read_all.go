package store

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func readAll(dir string) (map[string][]byte, error) {
	log.Println(dir)
	datas := make(map[string][]byte)
	err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}

			buf := make([]byte, info.Size())
			_, err = file.Read(buf)
			if err != nil {
				return err
			}

			datas[info.Name()] = buf
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return datas, nil
}
