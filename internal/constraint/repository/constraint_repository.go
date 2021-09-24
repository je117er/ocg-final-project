package constraint

import (
	"context"
	"database/sql"
	"github.com/je117er/ocg-final-project/internal/constraint"
	"github.com/je117er/ocg-final-project/internal/models"
	"github.com/je117er/ocg-final-project/internal/utils"
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
