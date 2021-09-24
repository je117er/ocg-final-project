package models

import (
	"database/sql"
	"time"
)

type Booking struct {
	ID                sql.NullInt64
	CustomerID        int64
	DateRegistered    sql.NullTime
	DateBooked        sql.NullTime
	TimePeriod        sql.NullInt64
	DosesCompleted    sql.NullInt64
	VaccineName       sql.NullString
	ExtendedInterval  sql.NullInt64
	LotNumber         sql.NullString
	ClinicName        sql.NullString
	Price             sql.NullString
	SentReminderEmail sql.NullInt64
	SessionCapacityID sql.NullInt64
	TotalBill         sql.NullString
	PaymentStatus     sql.NullBool
	StockItemID       sql.NullString
}

type VaccinationResponse struct {
	BookingID        int       `json:"bookingID"`
	Dose             int       `json:"dose"`
	VaccineName      string    `json:"vaccineName"`
	RegistrationTime time.Time `json:"registrationTime"`
	Completed        bool      `json:"completed"`
	VaccinationTime  time.Time `json:"vaccinationTime"`
	Clinic           string    `json:"clinic"`
}

type VaccinationRequest struct {
	BookingID int `json:"bookingID"`
	Completed int `json:"completed"`
}
