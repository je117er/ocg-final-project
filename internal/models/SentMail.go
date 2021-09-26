package models

import "time"

type SentMail struct {
	BookingID  int
	Email      string
	TimePeriod string
	ClinicName string
	Date       time.Time
	Name       string
}
