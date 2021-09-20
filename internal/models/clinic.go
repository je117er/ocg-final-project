package models

import "database/sql"

type Clinic struct {
	ID       string
	Name     sql.NullString
	Address  sql.NullString
	Contact  sql.NullString
	Location sql.NullString
	Status   sql.NullBool
}
