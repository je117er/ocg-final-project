package clinic

import (
	"context"
	"database/sql"
	"github.com/je117er/ocg-final-project/internal/clinic"
	"github.com/je117er/ocg-final-project/internal/models"
	"github.com/je117er/ocg-final-project/internal/utils"
)

type ClinicRepository struct {
	conn *sql.DB
}

func NewClinicRepository(conn *sql.DB) clinic.Repository {
	return &ClinicRepository{
		conn,
	}
}

var logger = utils.SugarLog()

func (repository *ClinicRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.Clinic, error) {
	rows, err := repository.conn.QueryContext(ctx, query, args...)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	defer rows.Close()

	result := make([]*models.Clinic, 0)
	for rows.Next() {
		clinic := new(models.Clinic)
		err := rows.Scan(&clinic.ID, &clinic.Name, &clinic.Address, &clinic.Contact, &clinic.Location, &clinic.Status)
		if err != nil {
			logger.Error(err)
			return nil, err
		}

		result = append(result, clinic)
	}

	return result, nil

}

func (repository *ClinicRepository) Fetch(ctx context.Context) ([]*models.Clinic, error) {
	query := "SELECT * FROM `clinic`"
	return repository.fetch(ctx, query)
}
