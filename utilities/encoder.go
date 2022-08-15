package utilities

import (
	"encoding/json"
	"errors"
	"io/fs"
	"os"
)

func EncodeToFile(input interface{}, filePath string) error {
	encodedData, err := json.Marshal(input)
	if err != nil {
		return err
	}

	err = os.WriteFile(filePath, encodedData, fs.ModePerm)
	return err
}

func DecodeFromFile(output interface{}, filePath string) error {
	encodedData, err := os.ReadFile(filePath)
	if errors.Is(err, os.ErrNotExist) {
		return nil
	}

	if err != nil {
		return err
	}

	err = json.Unmarshal(encodedData, &output)
	return err
}
