package store

import (
	"os"
)

type Store struct{}

func Open(path string) (*Store, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	f.Close()
	return &Store{}, nil
}
