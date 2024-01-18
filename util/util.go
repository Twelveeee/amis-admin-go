package util

import (
	"encoding/json"
	"io"
	"os"

	"gopkg.in/yaml.v2"
)

func ReadFileJson(filepath string) (map[string]interface{}, error) {
	var ret map[string]interface{}
	file, err := os.OpenFile(filepath, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	jsonBytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(jsonBytes, &ret)
	if err != nil {
		return nil, err
	}

	// return jsonString, nil
	return ret, nil
}

func LoadYaml(path string, v interface{}) error {
	file, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(file, v)
}
