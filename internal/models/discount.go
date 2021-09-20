package models

import "database/sql"

type Discount struct {
	ID          sql.NullInt32
	Code        sql.NullString
	UsageCount  sql.NullInt32
	PriceRuleID int32
}
