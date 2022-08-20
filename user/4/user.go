package user

type User struct {
	Name     string
	Language string
}

func New(name string) User {
	return User{
		Name: name,
	}
}

var greeting = map[string]string{
	"Chinese": "你好",
	"Spanish": "¿Que tal?",
}

func (u User) Greeting() string {
	return greeting[u.Language]
}
