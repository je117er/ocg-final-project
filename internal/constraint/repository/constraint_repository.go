package constraint

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/je117er/ocg-final-project/internal/constraint"
	"github.com/je117er/ocg-final-project/internal/models"
	"github.com/je117er/ocg-final-project/internal/utils"
	"strings"
)

type ConstraintRepository struct {
	conn *sql.DB
}

func NewConstraintRepository(conn *sql.DB) constraint.Repository {
	return &ConstraintRepository{
		conn,
	}
}

var logger = utils.SugarLog()

func (repository *ConstraintRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.Constraint, error) {
	rows, err := repository.conn.QueryContext(ctx, query, args...)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	defer rows.Close()

	result := make([]*models.Constraint, 0)
	for rows.Next() {
		cst := new(models.Constraint)
		err := rows.Scan(&cst.ID, &cst.Code, &cst.Description, &cst.VaccineEligible, &cst.Recommendation)
		if err != nil {
			logger.Error(err)
			return nil, err
		}

		result = append(result, cst)
	}

	return result, nil

}

func (repository *ConstraintRepository) Fetch(ctx context.Context) ([]*models.Constraint, error) {
	query := "SELECT * FROM `constraint`"
	return repository.fetch(ctx, query)
}

func (repository *ConstraintRepository) SaveConstraint(ctx context.Context, conditions []models.ConditionOrderRequest, customerID int) (*sql.Tx, error) {
	tx, err := repository.conn.BeginTx(ctx, nil)
	if err != nil {
		logger.Error(err)
		return tx, err
	}

	var valueStrs []string
	var valueArgs []interface{}

	for _, v := range conditions {
		valueStrs = append(valueStrs, "(?, ?, ?, ?)")
		valueArgs = append(valueArgs, v.Code, v.Description, v.ConditionStatus, 1)
	}
	query := fmt.Sprintf("INSERT INTO medical_condition (code, description, condition_status, customer_id)  VALUES %s",
		strings.Join(valueStrs, ","))

	_, err = tx.ExecContext(ctx, query, valueArgs...)
	if err != nil {
		tx.Rollback()
		return tx, err
	}

	return tx, nil
}
