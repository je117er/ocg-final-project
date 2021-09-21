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

type ClinicRequest struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	Contact  string `json:"contact"`
	Location string `json:"location"`
	Status   bool   `json:"status"`
}

type ClinicResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	Contact  string `json:"contact"`
	Location string `json:"location"`
	Status   bool   `json:"status"`
}

func (clinic Clinic) Entity2Response() *ClinicResponse {
	return &ClinicResponse{
		ID:       clinic.ID,
		Name:     clinic.Name.String,
		Address:  clinic.Address.String,
		Contact:  clinic.Contact.String,
		Location: clinic.Location.String,
		Status:   clinic.Status.Bool,
	}
}
