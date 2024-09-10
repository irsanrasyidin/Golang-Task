package repository

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Usecase3Repository interface {
	Create(filename string, fileData []byte) error
	List() ([]string, error)
	GetByID(filename string) ([]byte, error)
	Delete(filename string) error
}

type usecase3Repository struct {
}

func (e *usecase3Repository) Create(filename string, fileData []byte) error {
	// Create directory if it does not exist
	if err := os.MkdirAll("./uploads", os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directory: %v", err)
	}

	// Save file to the uploads directory
	filePath := filepath.Join("./uploads", filename)
	if err := ioutil.WriteFile(filePath, fileData, 0644); err != nil {
		return fmt.Errorf("failed to save file: %v", err)
	}

	return nil
}

func (e *usecase3Repository) List() ([]string, error) {
	files, err := filepath.Glob("./uploads/*")
	if err != nil {
		return nil, fmt.Errorf("failed to list files: %v", err)
	}

	var filenames []string
	for _, file := range files {
		filenames = append(filenames, filepath.Base(file))
	}

	return filenames, nil
}

func (e *usecase3Repository) GetByID(filename string) ([]byte, error) {
	filePath := filepath.Join("./uploads", filename)

	// Check if the file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return nil, fmt.Errorf("file not found: %v", err)
	}

	// Read file data
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %v", err)
	}

	return data, nil
}

func (e *usecase3Repository) Delete(filename string) error {
	filePath := filepath.Join("./uploads", filename)

	// Delete the file
	if err := os.Remove(filePath); err != nil {
		return fmt.Errorf("failed to delete file: %v", err)
	}

	return nil
}

func NewU3Repository() Usecase3Repository {
	return &usecase3Repository{}
}
