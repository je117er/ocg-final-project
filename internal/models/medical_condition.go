package models

import "database/sql"

type MedicalCondition struct {
	ID              sql.NullInt32
	Code            string
	Description     sql.NullString
	ConditionStatus sql.NullBool
	CustomerID      int32
}

type MedicalConditionResponse struct {
	ID              int32
	Code            string
	Description     string
	ConditionStatus bool
}

func (condition MedicalCondition) ModelToResponse() *MedicalConditionResponse {
	return &MedicalConditionResponse{
		ID:              condition.ID.Int32,
		Code:            condition.Code,
		Description:     condition.Description.String,
		ConditionStatus: condition.ConditionStatus.Bool,
	}
}
