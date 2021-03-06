package models

import (
	"database/sql"
	"time"
)

type SessionCapacity struct {
	ID          int
	ClinicID    sql.NullString
	Capacity    sql.NullInt64
	Type        sql.NullInt64
	Status      sql.NullBool
	CurrentDate sql.NullTime
	SlotLeft    sql.NullInt64
}

type SessionByClinic struct {
	ClinicID sql.NullString
	SessionID int
	TimePeriod sql.NullInt64
	CurrentDate sql.NullTime
}

type SessionByClinicResponse struct {
	ClinicID string `json:"clinic_id"`
	SessionID int `json:"session_id"`
	TimePeriod int `json:"time_period"`
	CurrentDate time.Time `json:"current_date"`
}

func (s SessionByClinic) ModelToResponse() *SessionByClinicResponse {
	return &SessionByClinicResponse{
		ClinicID: s.ClinicID.String,
		SessionID: s.SessionID,
		TimePeriod: int(s.TimePeriod.Int64),
		CurrentDate: s.CurrentDate.Time,
	}
}

type SessionCapacityResponse struct {
	ID          int       `json:"id"`
	ClinicID    string    `json:"clinicID"`
	Capacity    int       `json:"capacity"`
	Type        int       `json:"type"`
	Status      bool      `json:"status"`
	CurrentDate time.Time `json:"currentDate"`
	SlotLeft    int       `json:"slotLeft"`
}

type SessionCapacityRequest struct {
	ID             int    `json:"id"`
	ClinicID       string `json:"clinicID"`
	Capacity       int    `json:"capacity"`
	Type           int    `json:"type"`
	Status         bool   `json:"status"`
	CurrentDateStr string `json:"currentDateStr"`
	SlotLeft       int    `json:"slotLeft"`
}

func (session SessionCapacity) Entity2Response() *SessionCapacityResponse {
	return &SessionCapacityResponse{
		ID:          session.ID,
		ClinicID:    session.ClinicID.String,
		Capacity:    int(session.Capacity.Int64),
		Type:        int(session.Type.Int64),
		Status:      session.Status.Bool,
		CurrentDate: session.CurrentDate.Time,
		SlotLeft:    int(session.SlotLeft.Int64),
	}
}
