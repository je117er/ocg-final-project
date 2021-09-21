package models

import (
	"database/sql"
	"time"
)

type Product struct {
	ID                    string
	Name                  sql.NullString
	Price                 sql.NullString
	Vendor                sql.NullString
	VaccineType           sql.NullString
	AuthorizedAges        sql.NullInt32
	Dose                  sql.NullString
	AntigenNature         sql.NullString
	RouteOfAdministration sql.NullString
	StorageRequirements   sql.NullString
	AvailableFormats      sql.NullString
	Diluent               sql.NullString
	Adjuvant              sql.NullString
	AlternateName         sql.NullString
	MinimumInterval       sql.NullInt32
	ImmunizationSchedule  sql.NullInt32
	AuthorizedInterval    sql.NullInt32
	ExtendedInterval      sql.NullInt32
	Background            sql.NullString
	RegulatoryActions     sql.NullString
	SafetyStatus          sql.NullString
	AuthorizationStatus   sql.NullBool
	Trials                sql.NullString
	Distribution          sql.NullString
	Funding               sql.NullString
	Slug                  sql.NullString
	Image                 sql.NullString
	LotNumber             sql.NullString
	ExpiryDate            sql.NullTime
}

type ProductResponse struct {
	ID                    string
	Name                  string
	Price                 string
	Vendor                string
	VaccineType           string
	AuthorizedAges        int32
	Dose                  string
	AntigenNature         string
	RouteOfAdministration string
	StorageRequirements   string
	AvailableFormats      string
	Diluent               string
	Adjuvant              string
	AlternateName         string
	MinimumInterval       int32
	ImmunizationSchedule  int32
	AuthorizedInterval    int32
	ExtendedInterval      int32
	Background            string
	RegulatoryActions     string
	SafetyStatus          string
	AuthorizationStatus   bool
	Trials                string
	Distribution          string
	Funding               string
	Slug                  string
	Image                 string
	LotNumber             string
	ExpiryDate            time.Time
}

func (p Product) ModelToResponse() *ProductResponse {
	return &ProductResponse{
		ID:                    p.ID,
		Name:                  p.Name.String,
		Price:                 p.Price.String,
		Vendor:                p.Vendor.String,
		VaccineType:           p.VaccineType.String,
		AuthorizedAges:        p.AuthorizedAges.Int32,
		Dose:                  p.Dose.String,
		AntigenNature:         p.AntigenNature.String,
		RouteOfAdministration: p.RouteOfAdministration.String,
		StorageRequirements:   p.StorageRequirements.String,
		AvailableFormats:      p.AvailableFormats.String,
		Diluent:               p.Diluent.String,
		Adjuvant:              p.Adjuvant.String,
		AlternateName:         p.AlternateName.String,
		MinimumInterval:       p.MinimumInterval.Int32,
		ImmunizationSchedule:  p.ImmunizationSchedule.Int32,
		AuthorizedInterval:    p.AuthorizedInterval.Int32,
		ExtendedInterval:      p.ExtendedInterval.Int32,
		Background:            p.Background.String,
		RegulatoryActions:     p.RegulatoryActions.String,
		SafetyStatus:          p.SafetyStatus.String,
		AuthorizationStatus:   p.AuthorizationStatus.Bool,
		Trials:                p.Trials.String,
		Distribution:          p.Distribution.String,
		Funding:               p.Funding.String,
		Slug:                  p.Slug.String,
		Image:                 p.Image.String,
		LotNumber:             p.LotNumber.String,
		ExpiryDate:            p.ExpiryDate.Time,
	}
}
