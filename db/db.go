package db

import (
	"database/sql"
	"errors"
	"log"
	"os"
	"path/filepath"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type PassManager struct {
	db   *sql.DB
	path string
}

func getPath(dbName string) (string, error) {
	if dbName == "" {
		dbName = "default.db"
	}

	homedir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	d := filepath.Join(homedir, ".pass")
	if _, err := os.Stat(d); os.IsNotExist(err) {
		err := os.Mkdir(d, 0700)
		if err != nil {
			return "", err
		}
	}

	path := filepath.Join(d, dbName)
	return path, nil
}

func Open(dbName string) (*PassManager, error) {
	dbPath, err := getPath(dbName)
	if err != nil {
		return nil, errors.New("Unable to open database.")
	}

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}
	pm := PassManager{db, dbPath}
	return &pm, nil
}

func (pm *PassManager) Init() error {
	const createTables string = `
		CREATE TABLE IF NOT EXISTS accounts (
			id				TEXT NOT NULL PRIMARY KEY,
			name			TEXT NOT NULL UNIQUE,
			username	TEXT NULL,
			created		DATETIME NOT NULL,
		  updated		DATETIME NOT NULL
		);

		CREATE TABLE IF NOT EXISTS passwords (
			accountId TEXT NOT NULL REFERENCES accounts(id),
			pass			TEXT NOT NULL,
			created		DATETIME NOT NULL
		);
		`
	if _, err := pm.db.Exec(createTables); err != nil {
		return err
	}

	return nil
}

func (pm *PassManager) AddAccount(a *Account) (int, error) {
	const stmt string = "INSERT INTO accounts(id, name, username, created, updated) VALUES(?, ?, ?, ?, ?)"
	result, err := pm.db.Exec(stmt, a.Id, a.Name, a.Username, time.Now(), time.Now())

	if err != nil {
		return 0, err
	}
	lastId, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}
	return int(lastId), nil
}

func (pm *PassManager) GetAccountByName(name string) (Account, error) {
	const stmt string = "SELECT * FROM accounts WHERE name = ?"

	var account Account
	if err := pm.db.QueryRow(stmt, name).Scan(
		&account.Id,
		&account.Name,
		&account.Username,
		&account.Created,
		&account.Updated,
	); err != nil {
		log.Fatal(err)
	}
	return account, nil
}

func (pm *PassManager) GeneratePassword(a *Account) error {
	const stmt string = "INSERT INTO passwords(accountId, pass, created) VALUES (?, ?, ?)"

	password := "new_generated_password_here"

	_, err := pm.db.Exec(stmt, 100, password, time.Now())

	if err != nil {
		return err
	}

	return nil
}
