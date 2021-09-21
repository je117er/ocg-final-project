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

func (repository *ClinicRepository) getOne(ctx context.Context, query string, args ...interface{}) (*models.Clinic, error) {
	stmt, err := repository.conn.PrepareContext(ctx, query)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	row := stmt.QueryRowContext(ctx, args...)
	clinic := &models.Clinic{}
	err = row.Scan(&clinic.ID, &clinic.Name, &clinic.Address, &clinic.Contact, &clinic.Location, &clinic.Status)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return clinic, nil
}

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

func (repository *ClinicRepository) GetByID(ctx context.Context, id string) (*models.Clinic, error) {
	query := "SELECT * FROM `clinic` WHERE id = ?"
	return repository.getOne(ctx, query, id)
}

func (repository *ClinicRepository) Update(ctx context.Context, clinic models.ClinicRequest) (*models.Clinic, error) {
	query := "UPDATE `clinic` set name = ?, address = ?, contact = ?, location = ?, status = ? WHERE id = ?"
	stmt, err := repository.conn.PrepareContext(ctx, query)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	res, err := stmt.ExecContext(ctx, clinic.Name, clinic.Address, clinic.Contact, clinic.Location, clinic.Status, clinic.ID)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	_, err = res.RowsAffected()
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	data, err := repository.GetByID(ctx, clinic.ID)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return data, nil
}
