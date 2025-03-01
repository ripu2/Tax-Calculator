package fileops

import (
	"encoding/json"
	"errors"
	"os"
)

func OpenFile(path string) (*os.File, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return file, nil //... rest of the program
}

func WriteToJson(path string, data interface{}) error {
	// Convert struct to pretty JSON
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return errors.New(err.Error())
	}

	// Explicitly create the file
	file, err := os.Create(path)
	if err != nil {
		return errors.New(err.Error())
	}
	defer file.Close() // Ensure file is closed after function exits

	// Write JSON to the file
	_, err = file.Write(jsonData)
	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}
