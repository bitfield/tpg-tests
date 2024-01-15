package user

import "sync"

type User struct {
	Name string
}

var (
	m     = new(sync.Mutex)
	users = map[string]*User{}
)

func Create(name string) {
	m.Lock()
	defer m.Unlock()
	users[name] = &User{
		Name: name,
	}
}

func Delete(name string) {
	m.Lock()
	defer m.Unlock()
	delete(users, name)
}

func Exists(name string) bool {
	m.Lock()
	defer m.Unlock()
	_, ok := users[name]
	return ok
}
