package config

import (
	"encoding/json"
	"io/ioutil"
)

func loadJSONFromFile(fileName string) (fileValueMap map[string]interface{}, err error) {
	fileValueMap = make(map[string]interface{})

	fileContent, err := ioutil.ReadFile(fileName)

	if err != nil {
		return
	}

	err = json.Unmarshal(fileContent, &fileValueMap)
	if err != nil {
		return fileValueMap, nil
	}
	return
}

func LoadJSON(fileName string) (map[string]interface{}, error) {
	return loadJSONFromFile(fileName)
}
