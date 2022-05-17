package utils

import (
	"fmt"
	"io/fs"
	"log"
	"path/filepath"

	"github.com/xeipuuv/gojsonschema"
)

var tmpDir = "../tmp"
var schemaDir = "../schema"

func InitSchemas() {
	LoadersInstance = make(map[string]gojsonschema.JSONLoader)
	readAllSchemas()
}

var LoadersInstance Loaders

type Loaders map[string]gojsonschema.JSONLoader

func LoadSchema(name string) error {
	abs, err := filepath.Abs(fmt.Sprintf("../schema/%v", name))
	if err != nil {
		return fmt.Errorf("cannot get absolute path: %w", err)
	}

	LoadersInstance[name] = gojsonschema.NewReferenceLoader(canonicalFormat(abs))

	return nil
}

func ValidateData(validator string, dataname string) (bool, error) {
	abs, err := filepath.Abs(dataname)
	if err != nil {
		return false, fmt.Errorf("cannot get absolute path: %w", err)
	}
	documentLoader := gojsonschema.NewReferenceLoader(canonicalFormat(abs))

	if _, exist := LoadersInstance[validator]; !exist {
		return false, fmt.Errorf("this file_type does not exist")
	} else {
		result, err := gojsonschema.Validate(LoadersInstance[validator], documentLoader)
		if err != nil {
			return false, err
		}

		return result.Valid(), nil
	}

}

func readAllSchemas() error {
	schemaNames := make([]string, 0)
	err := filepath.Walk(schemaDir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		schemaNames = append(schemaNames, info.Name())
		return nil
	})

	if err != nil {
		return err
	}

	for _, name := range schemaNames {
		abs, err := filepath.Abs(name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(name)
		LoadersInstance[name] = gojsonschema.NewReferenceLoader(canonicalFormat(abs))
		log.Println("loaded schema", name)
	}

	return nil
}

func canonicalFormat(name string) string {
	return fmt.Sprintf("file://%v", name)
}
