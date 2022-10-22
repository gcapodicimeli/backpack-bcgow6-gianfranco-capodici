package store

import (
	"encoding/json"
	"os"
)

type Store interface {
	Read(data interface{}) error
	Write(data interface{}) error
}

type Type string

const (
	FileType Type = "file"
	MonoType Type = "mongo" //Es un ejemplo si tuvieramos otra conexion a la bd
)

type FileStore struct {
	FilePath string
}

func New(store Type, fileName string) Store {
	switch store {
	case FileType:
		return &FileStore{FilePath: fileName}
	default:
		return nil
	}
}

func (fs *FileStore) Write(data interface{}) error {
	fileData, err := json.MarshalIndent(data, "", "")
	if err != nil {
		return err
	}
	return os.WriteFile(fs.FilePath, fileData, 0644)
}

func (fs *FileStore) Read(data interface{}) error {
	fileData, err := os.ReadFile(fs.FilePath)
	if err != nil {
		return err
	}
	return json.Unmarshal(fileData, &data)
}
