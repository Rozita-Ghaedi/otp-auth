package models

import "time"

type User struct {
	ID         string    `db:"id" json:"id"`
	Identifier string    `db:"identifier" json:"identifier"`
	Verified   bool      `db:"verified" json:"verified"`
	CreatedAt  time.Time `db:"created_at" json:"createdAt"`
	LastLogin  time.Time `db:"last_login" json:"lastLogin"`
}
