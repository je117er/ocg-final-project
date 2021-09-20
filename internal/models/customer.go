package models

import "database/sql"

type Customer struct {
	ID              int32
	Email           sql.NullString
	Name            sql.NullString
	Address         sql.NullString
	Gender          sql.NullString
	Dob             sql.NullTime
	PhoneNumber     sql.NullString
	InsuranceNumber sql.NullString
	City            sql.NullString
	District        sql.NullString
	Commune         sql.NullString
	Ethnicity       sql.NullString
	Nationality     sql.NullString
}
