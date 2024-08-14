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
	"French":  "Bonjour",
	"Spanish": "¿Que tal?",
}

func (u User) Greeting() string {
	return greeting[u.Language]
}
