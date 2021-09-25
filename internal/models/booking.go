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

type CustomerOrderRequest struct {
	CustomerName    string    `json:"customer_name"`
	Email           string    `json:"email"`
	Address         string    `json:"address"`
	Gender          string    `json:"gender"`
	DoB             time.Time `json:"dob"`
	PhoneNumber     string    `json:"phone_number"`
	InsuranceNumber string    `json:"insurance_number"`
	City            string    `json:"city"`
	District        string    `json:"district"`
	Commune         string    `json:"commune"`
	Ethnicity       string    `json:"ethnicity"`
	Nationality     string    `json:"nationality"`
}

type ConditionOrderRequest struct {
	Code            string `json:"code"`
	Description     string `json:"description"`
	ConditionStatus bool   `json:"condition_status"`
}

type OrderRequest struct {
	CustomerOrderRequest  CustomerOrderRequest
	ConditionOrderRequest []ConditionOrderRequest
	DateRegistered        time.Time `json:"date_registered"`
	DateBooked            time.Time `json:"date_booked"`
	TimePeriod            int64     `json:"time_period"`
	DosesCompleted        int64     `json:"doses_completed"`
	VaccineName           string    `json:"vaccine_name"`
	AuthorizedInterval    int       `json:"authorized_interval"`
	LotNumber             string    `json:"lot_number"`
	ClinicName            string    `json:"clinic_name"`
	Price                 int       `json:"price"`
	SentReminderEmail     int64     `json:"sent_reminder_email"`
	SessionCapacityID     int64     `json:"session_capacity_id"`
	TotalBill             int       `json:"total_bill"`
	StockItemID           string    `json:"stock_item_id"`
}

type CreateCheckoutSessionRequest struct {
	CustomerID int    `json:"customer_id"`
	TotalBill  int    `json:"total_bill"`
	Quantity   int    `json:"quantity"`
	SuccessURL string `json:"success_url"`
	CancelURL  string `json:"cancel_url"`
}
