package condition

import (
	"context"
	"database/sql"
	"github.com/je117er/ocg-final-project/internal/condition"
	"github.com/je117er/ocg-final-project/internal/models"
	"github.com/je117er/ocg-final-project/internal/utils"
)

type ConditionRepository struct {
	conn *sql.DB
}

func NewConditionRepository(conn *sql.DB) condition.Repository {
	return &ConditionRepository{
		conn,
	}
}

var logger = utils.SugarLog()

func (repository *ConditionRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.MedicalCondition, error) {
	rows, err := repository.conn.QueryContext(ctx, query, args...)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	defer rows.Close()

	result := make([]*models.MedicalCondition, 0)
	for rows.Next() {
		cdt := new(models.MedicalCondition)
		err := rows.Scan(&cdt.ID, &cdt.Code, &cdt.Description, &cdt.ConditionStatus)
		if err != nil {
			logger.Error(err)
			return nil, err
		}

		result = append(result, cdt)
	}

	return result, nil

}

func (repository *ConditionRepository) ByCustomerID(ctx context.Context, id int) ([]*models.MedicalCondition, error) {
	query := "SELECT ID, code, description, condition_status FROM `medical_history` WHERE customer_id = ?"
	return repository.fetch(ctx, query, id)
}
