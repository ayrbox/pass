package db

import "time"

type Account struct {
	Id       string
	Name     string
	Username string
	Created  time.Time
	Update   time.Time
}
