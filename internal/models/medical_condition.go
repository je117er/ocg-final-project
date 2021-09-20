package models

import "database/sql"

type MedicalCondition struct {
	ID              sql.NullInt32
	Code            string
	Description     sql.NullString
	ConditionStatus sql.NullBool
	CustomerID      int32
}
