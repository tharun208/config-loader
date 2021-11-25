package config

import (
	"encoding/json"
)

// Decoder for json
var JSONDecoder Decoder = func(data []byte) (map[string]interface{}, error) {
	return loadJSON(data)
}

func loadJSON(fileContent []byte) (fileValueMap map[string]interface{}, err error) {
	fileValueMap = make(map[string]interface{})

	err = json.Unmarshal(fileContent, &fileValueMap)
	if err != nil {
		return fileValueMap, err
	}
	return
}
