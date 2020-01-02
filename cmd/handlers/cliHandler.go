package handlers

import (
	"errors"
	"os"
	"path/filepath"
)

func ErrorHandler(configFile []string) (string, error) {
	if len(configFile) != 1 {
		return "", errors.New("as an argument you should only include path to flatters.json")
	}

	filePath := configFile[0]
	if err := fileExists(filePath); err != nil {
		return "", err
	}

	return filePath, nil
}

func fileExists(filename string) error {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return errors.New("file is not exists")
	}
	if info.IsDir() {
		return errors.New("path's end cannot be a directory")
	}
	if fileExt := filepath.Ext(info.Name()); fileExt != ".json" {
		return errors.New("given file is not json")
	}
	return nil
}
