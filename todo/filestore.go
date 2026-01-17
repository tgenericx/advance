package todo

import (
	"encoding/json"
	"os"
)

type FileStore struct {
	path string
}

func NewFileStore(path string) *FileStore {
	return &FileStore{path: path}
}

func (fs *FileStore) Load() ([]Item, error) {
	data, err := os.ReadFile(fs.path)
	if err != nil {
		if os.IsNotExist(err) {
			return []Item{}, nil
		}
		return nil, err
	}

	var items []Item
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, err
	}

	return items, nil
}

func (fs *FileStore) Save(items []Item) error {
	data, err := json.Marshal(items)
	if err != nil {
		return err
	}

	return os.WriteFile(fs.path, data, 0o644)
}
