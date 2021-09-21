package models

import (
	"database/sql"
	"time"
)

type Customer struct {
	ID              int
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
	ID              int       `json:"id"`
	Email           string    `json:"email"`
	Name            string    `json:"name"`
	Address         string    `json:"address"`
	Gender          string    `json:"gender"`
	Dob             time.Time `json:"dob"`
	PhoneNumber     string    `json:"phoneNumber"`
	InsuranceNumber string    `json:"insuranceNumber"`
	City            string    `json:"city"`
	District        string    `json:"district"`
	Commune         string    `json:"commune"`
	Ethnicity       string    `json:"ethnicity"`
	Nationality     string    `json:"nationality"`
}

type CustomerRequest struct {
	ID              int       `json:"id"`
	Email           string    `json:"email"`
	Name            string    `json:"name"`
	Address         string    `json:"address"`
	Gender          string    `json:"gender"`
	Dob             time.Time `json:"dob"`
	PhoneNumber     string    `json:"phoneNumber"`
	InsuranceNumber string    `json:"insuranceNumber"`
	City            string    `json:"city"`
	District        string    `json:"district"`
	Commune         string    `json:"commune"`
	Ethnicity       string    `json:"ethnicity"`
	Nationality     string    `json:"nationality"`
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
