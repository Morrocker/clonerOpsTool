package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	st "github.com/clonerOpsTool/pkg/structs"
)

// UploadStorageConf asdfasf asdfa sdf
func UploadStorageConf(filepath string) (st.StorageConfig, error) {
	var data st.StorageConfig
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

// UploadInstructions asdfa asdfas
func UploadInstructions(filepath string) (st.Instructions, error) {
	var data st.Instructions
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
func WriteJSON(name string, data st.StorageConfig) error {
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
