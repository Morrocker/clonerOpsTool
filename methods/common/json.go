package common

import (
	"encoding/json"
	"io/ioutil"
)

// UploadJSON asdfasf asdfa sdf
func UploadJSON(filepath string) (StorageConfig, error) {
	var data StorageConfig
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
