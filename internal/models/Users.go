package models

import "time"

type User struct {
	ID       string
	Username string
	CreateAT *time.Time
	UpdateAT *time.Time
}
