package models

import (
	"database/sql"
	"time"
)

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

type CustomerResponse struct {
	ID              int32
	Email           string
	Name            string
	Address         string
	Gender          string
	Dob             time.Time
	PhoneNumber     string
	InsuranceNumber string
	City            string
	District        string
	Commune         string
	Ethnicity       string
	Nationality     string
}

func (customer Customer) Entity2Response() *CustomerResponse {
	return &CustomerResponse{
		customer.ID,
		customer.Email.String,
		customer.Name.String,
		customer.Address.String,
		customer.Gender.String,
		customer.Dob.Time,
		customer.PhoneNumber.String,
		customer.InsuranceNumber.String,
		customer.City.String,
		customer.District.String,
		customer.Commune.String,
		customer.Ethnicity.String,
		customer.Nationality.String,
	}
}
