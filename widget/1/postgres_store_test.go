package widget_test

import (
	"testing"

	"widget"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/go-cmp/cmp"
)

func TestPostgresStore_Retrieve(t *testing.T) {
	t.Parallel()
	ps := fakePostgresStore(t)
	want := widget.Widget{
		ID:   "widget01",
		Name: "Acme Giant Rubber Band",
	}
	got, err := ps.Retrieve("widget01")
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func fakePostgresStore(t *testing.T) widget.PostgresStore {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		db.Close()
	})
	query := "SELECT id, name FROM widgets"
	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow("widget01", "Acme Giant Rubber Band")
	mock.ExpectQuery(query).WillReturnRows(rows)
	return widget.PostgresStore{
		DB: db,
	}
}
