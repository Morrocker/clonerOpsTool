package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	ed "github.com/clonerOpsTool/pkg/confeditor"
)

// UploadStorageConf asdfasf asdfa sdf
func UploadStorageConf(filepath string) (ed.StorageConfig, error) {
	var sc ed.StorageConfig
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		return sc, err
	}

	err = json.Unmarshal([]byte(file), &sc)
	if err != nil {
		return sc, err
	}
	// spew.Dump(data)
	return sc, nil

}

// UploadInstructions asdfa asdfas
func UploadInstructions(filepath string) (ed.Instructions, error) {
	var data ed.Instructions
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		return data, err
	}

	err = json.Unmarshal([]byte(file), &data)
	if err != nil {
		return data, err
	}
	// spew.Dump(data)
	return data, nil

}

// WriteJSON asasdfadf asfdasdf asdf asdfsdfsafd
func WriteJSON(name string, data interface{}) error {
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(name, file, 0644)
	if err != nil {
		return err

	}
	fmt.Printf("JSON data writen to file %s\n", name)
	return nil
}
