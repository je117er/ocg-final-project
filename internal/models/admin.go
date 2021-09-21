package models

import (
	"database/sql"
)

type Admin struct {
	ID             int32
	Username       sql.NullString
	HashedPassword sql.NullString
}

type LoginRequest struct {
	Username string
	Password string
}

type LoginResponse struct {
	Status   bool
	Message  string
	Username string
	Token    string
}
