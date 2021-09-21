package models

import "database/sql"

type Constraint struct {
	ID              int
	Code            string
	Description     sql.NullString
	VaccineEligible sql.NullBool
	NeedPrecaution  sql.NullBool
}

type ConstraintResponse struct {
	ID              int
	Code            string
	Description     string
	VaccineEligible bool
	NeedPrecaution  bool
}

func (constraint Constraint) Entity2Response() *ConstraintResponse {
	return &ConstraintResponse{
		ID:              constraint.ID,
		Code:            constraint.Code,
		Description:     constraint.Description.String,
		VaccineEligible: constraint.VaccineEligible.Bool,
		NeedPrecaution:  constraint.NeedPrecaution.Bool,
	}
}
