package models

import (
	"database/sql"
	"strconv"
)

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

type ClinicByVaccine struct {
	ClinicName         sql.NullString
	StockItemID        sql.NullString
	StockName          sql.NullString
	AuthorizedInterval sql.NullInt32
	LotNumber          sql.NullString
	Price              sql.NullString
}

type ClinicByVaccineResponse struct {
	ClinicName         string `json:"clinic_name"`
	StockItemID        string `json:"stock_item_id"`
	StockName          string `json:"stock_name"`
	AuthorizedInterval int    `json:"authorized_interval"`
	LotNumber          string `json:"lot_number"`
	Price              int    `json:"price"`
}

func (c ClinicByVaccine) ModelToResponse() *ClinicByVaccineResponse {
	price, _ := strconv.Atoi(c.Price.String)
	return &ClinicByVaccineResponse{
		ClinicName:         c.ClinicName.String,
		StockName:          c.StockName.String,
		StockItemID:        c.StockItemID.String,
		AuthorizedInterval: int(c.AuthorizedInterval.Int32),
		LotNumber:          c.LotNumber.String,
		Price:              price,
	}
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
