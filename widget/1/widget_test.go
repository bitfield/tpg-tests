package widget_test

import (
	"sync"
	"testing"

	"widget"

	"github.com/google/go-cmp/cmp"
)

func TestCreate_GivesNoErrorForValidWidget(t *testing.T) {
	s := newMapStore()
	w := widget.Widget{
		ID:   "widget01",
		Name: "Acme Giant Rubber Band",
	}
	wantID := "widget01"
	gotID, err := widget.Create(s, w)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if wantID != gotID {
		t.Error(cmp.Diff(wantID, gotID))
	}
}

type mapStore struct {
	m    *sync.Mutex
	data map[string]widget.Widget
}

func newMapStore() *mapStore {
	return &mapStore{
		m:    new(sync.Mutex),
		data: map[string]widget.Widget{},
	}
}

func (ms *mapStore) Store(w widget.Widget) (string, error) {
	ms.m.Lock()
	defer ms.m.Unlock()
	ms.data[w.ID] = w
	return w.ID, nil
}
