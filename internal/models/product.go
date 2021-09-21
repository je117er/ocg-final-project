package models

import (
	"database/sql"
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


