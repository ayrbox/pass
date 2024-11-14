package db

import (
	"database/sql"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

type PassManager struct {
	DB   *sql.DB
	Path string
}

func Init(path string, dbName string) (*PassManager, error) {
	f := filepath.Join(path, dbName)
	db, err := sql.Open("sqlite3", f)

	if err != nil {
		return nil, err
	}

	const createTables string = `
		CREATE TABLE IF NOT EXISTS accounts (
			id			INTEGER NOT NULL PRIMARY KEY,
			name		TEXT NOT NULL,
			created DATETIME NOT NULL,
		  updated DATETIME NOT NULL
		);`
	if _, err := db.Exec(createTables); err != nil {
		return nil, err
	}

	pm := PassManager{db, f}

	return &pm, nil
}
