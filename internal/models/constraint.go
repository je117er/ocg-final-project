package models

import "database/sql"

type Constraint struct {
	ID              int
	Code            string
	Description     sql.NullString
	VaccineEligible sql.NullInt32
	Recommendation  string
}

type ConstraintResponse struct {
	ID              int
	Code            string
	Description     string
	VaccineEligible int32
	Recommendation  string
}

func (constraint Constraint) Entity2Response() *ConstraintResponse {
	return &ConstraintResponse{
		ID:              constraint.ID,
		Code:            constraint.Code,
		Description:     constraint.Description.String,
		VaccineEligible: constraint.VaccineEligible.Int32,
		Recommendation:  constraint.Recommendation,
	}
}
