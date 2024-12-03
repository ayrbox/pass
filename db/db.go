package db

import (
	"database/sql"
	"errors"
	"fmt"
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

	connString := fmt.Sprintf("%v?_foreign_keys=on", dbPath)

	db, err := sql.Open("sqlite3", connString)
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
  accountId  TEXT NOT NULL,
  pass			 TEXT NOT NULL,
  created		 DATETIME NOT NULL,
  archived   DATETIME NULL,
  FOREIGN KEY(accountId) REFERENCES accounts(id)
);`

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

func (pm *PassManager) UpdateAccountName(a *Account, newAccountName string) error {
	const stmt string = "UPDATE accounts SET name = ? WHERE id = ?"
	_, err := pm.db.Exec(stmt, newAccountName, a.Id)
	if err != nil {
		return err
	}
	return nil
}

func (pm *PassManager) UpdateUsername(a *Account, newUsername string) error {
	const stmt string = "UPDATE accounts SET username = ? WHERE id = ?"
	_, err := pm.db.Exec(stmt, newUsername, a.Id)
	if err != nil {
		return err
	}
	return nil
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

func (pm *PassManager) UpdatePassword(a *Account, newPassword string) error {
	// TODO: use db transacation
	// TODO: generate random password with option of symbol and numbers

	// archive existing passwords
	const update_stmt string = "UPDATE passwords SET archived = ? WHERE accountId = ? AND archived IS NULL"
	_, err := pm.db.Exec(update_stmt, time.Now(), a.Id)
	if err != nil {
		return err
	}

	// insert new password
	const stmt string = "INSERT INTO passwords(accountId, pass, created) VALUES (?, ?, ?)"
	_, err = pm.db.Exec(stmt, a.Id, newPassword, time.Now())

	if err != nil {
		return err
	}

	return nil
}

func (pm *PassManager) GeneratePassword(a *Account) error {
	// TODO: use db transacation
	password, err := generatePassword(20, true, true)
	if err != nil {
		return err
	}

	// archive existing passwords
	const update_stmt string = "UPDATE passwords SET archived = ? WHERE accountId = ? AND archived IS NULL"
	_, err = pm.db.Exec(update_stmt, time.Now(), a.Id)
	if err != nil {
		return err
	}

	// insert new password
	const stmt string = "INSERT INTO passwords(accountId, pass, created) VALUES (?, ?, ?)"
	_, err = pm.db.Exec(stmt, a.Id, password, time.Now())

	if err != nil {
		return err
	}

	return nil
}
