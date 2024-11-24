package db

import (
	"database/sql"
	"path/filepath"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type PassManager struct {
	db   *sql.DB
	path string
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

func (pm *PassManager) addAccount(a *Account) (int, error) {
	const stmt string = "INSERT INTO accounts(id, name, created, updated) VALUES(?, ?, ?, ?)"
	result, err := pm.db.Exec(stmt, 100, a.name, time.Now(), time.Now())

	if err != nil {
		return 0, err
	}
	lastId, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}
	return int(lastId), nil
}

func (pm *PassManager) generatePassword(a *Account) (bool, error) {

	const stmt string = "INSERT INTO password(account_id, password, created) VALUES (?, ?, ?)"
	_, err := pm.db.Exec(stmt, 100, "newpassword", time.Now())

	if err != nil {
		return false, err
	}

	return true, nil
}
