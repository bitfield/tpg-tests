package store

import (
	"errors"
	"os"
)

var ErrUnopenable = errors.New("can't open store file")

type Store struct{}

func Open(path string) (*Store, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, ErrUnopenable // losing information about 'err'
	}
	f.Close()
	return &Store{}, nil
}
