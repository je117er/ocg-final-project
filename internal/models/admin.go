package models

import "database/sql"

type Admin struct {
	ID       int32
	Username sql.NullString
	Password sql.NullString
}
