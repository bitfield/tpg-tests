package kvstore

import "sync"

type Store struct {
	m    *sync.Mutex
	data map[string]string
}

func NewStore() *Store {
	return &Store{
		m:    new(sync.Mutex),
		data: map[string]string{},
	}
}

func (s *Store) Set(k, v string) {
	s.m.Lock()
	defer s.m.Unlock()
	s.data[k] = v
}

func (s *Store) Get(k string) string {
	s.m.Lock()
	defer s.m.Unlock()
	return s.data[k]
}
