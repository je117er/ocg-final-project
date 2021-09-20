package models

import "database/sql"

type Constraint struct {
	ID              int32
	Code            string
	Description     sql.NullString
	VaccineEligible sql.NullBool
	NeedPrecaution  sql.NullBool
}
