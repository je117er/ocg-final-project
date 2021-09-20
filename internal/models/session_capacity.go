package models

import "database/sql"

type SessionCapacity struct {
	ID          int32
	ClinicID    sql.NullString
	Capacity    sql.NullInt32
	Type        sql.NullInt32
	Status      sql.NullInt32
	CurrentDate sql.NullTime
}
