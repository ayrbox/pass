package db

import "time"

type Account struct {
	Id      int
	name    string
	created time.Time
	update  time.Time
}
