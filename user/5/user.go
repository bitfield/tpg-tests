package user5

import "sync"

type User struct {
	Name string
}

var (
	m     = new(sync.Mutex)
	users = map[string]*User{
		"Alice": {
			Name: "Alice",
		},
	}
)

func Create(name string) {}

func Exists(name string) bool {
	m.Lock()
	defer m.Unlock()
	_, ok := users[name]
	return ok
}
