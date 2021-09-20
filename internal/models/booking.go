package models

import "database/sql"

type Booking struct {
	ID                sql.NullInt32
	CustomerID        int32
	DateRegistered    sql.NullTime
	DateBooked        sql.NullTime
	TimePeriod        sql.NullInt32
	DosesCompleted    sql.NullInt32
	VaccineName       sql.NullString
	ExtendedInterval  sql.NullInt32
	LotNumber         sql.NullString
	ClinicName        sql.NullString
	Price             sql.NullString
	SentReminderEmail sql.NullInt32
	DailyCapacityID   sql.NullInt32
	TotalBill         sql.NullString
	PaymentStatus     sql.NullBool
	StockItemID       sql.NullString
	DiscountID        int32
}
