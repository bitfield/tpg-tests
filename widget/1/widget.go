package widget

type Widget struct {
	ID   string
	Name string
}

type Store interface {
	Store(Widget) (string, error)
}

func Create(s Store, w Widget) (ID string, err error) {
	ID, err = s.Store(w)
	if err != nil {
		return "", err
	}
	return ID, nil
}
