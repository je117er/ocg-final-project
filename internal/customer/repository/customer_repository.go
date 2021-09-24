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

func (repository *CustomerRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.Customer, error) {
	rows, err := repository.conn.QueryContext(ctx, query, args...)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	defer rows.Close()

	result := make([]*models.Customer, 0)
	for rows.Next() {
		customerModel := new(models.Customer)
		err := rows.Scan(&customerModel.ID, &customerModel.Email, &customerModel.Name, &customerModel.Address,
			&customerModel.Gender, &customerModel.Dob, &customerModel.PhoneNumber, &customerModel.InsuranceNumber,
			&customerModel.City, &customerModel.District, &customerModel.Commune, &customerModel.Ethnicity, &customerModel.Nationality)
		if err != nil {
			logger.Error(err)
			return nil, err
		}

		result = append(result, customerModel)
	}

	return result, nil

}

func (repository *CustomerRepository) GetAll(ctx context.Context) ([]*models.Customer, error) {
	query := `SELECT * FROM customer`
	result, err := repository.fetch(ctx, query)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (repository *CustomerRepository) GetByEmail(ctx context.Context, email string) (*models.Customer, error) {
	query := `SELECT c.id, c.email, c.name, c.address, c.gender, c.dob, c.phone_number, c.insurance_number, 
					c.city, c.district, c.commune, c.ethnicity, c.nationality 
				FROM customer AS c 
				WHERE c.email = ?`
	return repository.getOne(ctx, query, email)
}

func (repository *CustomerRepository) GetByID(ctx context.Context, id int) (*models.Customer, error) {
	query := `SELECT c.id, c.email, c.name, c.address, c.gender, c.dob, c.phone_number, c.insurance_number, 
					c.city, c.district, c.commune, c.ethnicity, c.nationality 
				FROM customer AS c 
				WHERE c.id = ?`
	return repository.getOne(ctx, query, id)
}

func (repository *CustomerRepository) Update(ctx context.Context, customer models.CustomerRequest) (*models.Customer, error) {
	query := `UPDATE customer SET email = ?, name = ?, address = ?, gender = ?, dob = ?, phone_number = ?,
				insurance_number = ?, city = ?, district = ?, commune = ?, ethnicity = ?, nationality = ?
				where id = ?`
	stmt, err := repository.conn.PrepareContext(ctx, query)
	if err != nil {
		return nil, nil
	}

	res, err := stmt.ExecContext(ctx, customer.Email, customer.Name, customer.Address, customer.Gender, customer.Dob,
		customer.PhoneNumber, customer.InsuranceNumber, customer.City, customer.District, customer.Commune, customer.Ethnicity,
		customer.Nationality, customer.ID)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	_, err = res.RowsAffected()
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	data, err := repository.GetByID(ctx, customer.ID)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return data, nil

}

func (repository *CustomerRepository) Save(ctx context.Context, customer models.CustomerRequest) (*models.Customer, error) {
	query := `INSERT INTO customer(email, name, address, gender, dob, phone_number, insurance_number, city, district, commune, ethnicity, nationality) 
				VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	stmt, err := repository.conn.PrepareContext(ctx, query)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	res, err := stmt.ExecContext(ctx, customer.Email, customer.Name, customer.Address, customer.Gender, customer.Dob,
		customer.PhoneNumber, customer.InsuranceNumber, customer.City, customer.District, customer.Commune, customer.Ethnicity,
		customer.Nationality)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	customer.ID = int(id)
	data, err := repository.GetByID(ctx, customer.ID)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return data, nil

}

func (repository *CustomerRepository) GetCertByID(ctx context.Context, id int) (*models.CustomerResponse, error) {
	return nil, nil
}

func (repository *CustomerRepository) GetAllByClinicID(ctx context.Context, id string) ([]*models.Customer, error) {
	query := `select *
			from customer
			where id in (
				select customer_id
				from booking
				where session_capacity_id in
					  (select id from session_capacity where clinic_id = ?)
			)`

	return repository.fetch(ctx, query, id)
}
