package repository

import (
	"context"
	"database/sql"
	"github.com/je117er/ocg-final-project/internal/customer"
	"github.com/je117er/ocg-final-project/internal/models"
	"github.com/je117er/ocg-final-project/internal/utils"
)

type CustomerRepository struct {
	conn *sql.DB
}

func NewCustomerRepository(conn *sql.DB) customer.Repository {
	return &CustomerRepository{
		conn,
	}
}

var logger = utils.SugarLog()

func (repository *CustomerRepository) getOne(ctx context.Context, query string, args ...interface{}) (*models.Customer, error) {
	stmt, err := repository.conn.PrepareContext(ctx, query)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	row := stmt.QueryRowContext(ctx, args...)
	customerModel := &models.Customer{}
	err = row.Scan(&customerModel.ID, &customerModel.Email, &customerModel.Name, &customerModel.Address,
		&customerModel.Gender, &customerModel.Dob, &customerModel.PhoneNumber, &customerModel.InsuranceNumber,
		&customerModel.City, &customerModel.District, &customerModel.Commune, &customerModel.Ethnicity, &customerModel.Nationality)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return customerModel, nil
}

func (repository *CustomerRepository) GetByEmail(ctx context.Context, email string) (*models.Customer, error) {
	query := `SELECT c.id, c.email, c.name, c.address, c.gender, c.dob, c.phone_number, c.insurance_number, 
					c.city, c.district, c.commune, c.ethnicity, c.nationality 
				FROM customer AS c 
				WHERE c.email = ?`
	return repository.getOne(ctx, query, email)
}
