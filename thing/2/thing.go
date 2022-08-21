package thing2

type Thing struct {
	X, Y, Z int
}

func NewThing(x, y, z int) (*Thing, error) {
	return &Thing{}, nil
}
