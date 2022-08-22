package widget

import (
	"context"
	"database/sql"
)

type PostgresStore struct {
	DB *sql.DB
}

func (p *PostgresStore) Store(w Widget) (ID string, err error) {
	// horrible SQL goes here
	// handle errors, etc
	return ID, nil
}

func (ps *PostgresStore) Retrieve(ID string) (Widget, error) {
	w := Widget{}
	ctx := context.Background()
	row := ps.DB.QueryRowContext(ctx, "SELECT id, name FROM widgets WHERE id = ?", ID)
	err := row.Scan(&w.ID, &w.Name)
	if err != nil {
		return Widget{}, err
	}
	return w, nil
}
