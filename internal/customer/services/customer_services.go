package services

import (
	"context"
	"github.com/je117er/ocg-final-project/internal/customer"
	"github.com/je117er/ocg-final-project/internal/models"
	"github.com/je117er/ocg-final-project/internal/utils"
	"github.com/pkg/errors"
)

type CustomerService struct {
	customerRepository customer.Repository
}

func NewCustomerService(customerRepository customer.Repository) customer.Service {
	return &CustomerService{
		customerRepository: customerRepository,
	}
}

var logger = utils.SugarLog()

func (customerService *CustomerService) GetByEmail(ctx context.Context, email string) (*models.CustomerResponse, error) {
	res, err := customerService.customerRepository.GetByEmail(ctx, email)
	if err != nil {
		return nil, errors.Errorf("Email %s not exist", email)
	}

	return res.Entity2Response(), nil
}

func (customerService *CustomerService) GetByID(ctx context.Context, id int) (*models.CustomerResponse, error) {
	res, err := customerService.customerRepository.GetByID(ctx, id)
	if err != nil {
		return nil, errors.Errorf("Customer %d doesn't exist", id)
	}

	return res.Entity2Response(), nil
}

func (customerService *CustomerService) GetAll(ctx context.Context) ([]*models.CustomerResponse, error) {
	res, err := customerService.customerRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	responses := make([]*models.CustomerResponse, 0)
	for _, v := range res {
		responses = append(responses, v.Entity2Response())
	}
	return responses, nil
}

func (customerService *CustomerService) UpdateCustomer(ctx context.Context, customerRequest models.CustomerRequest) (*models.CustomerResponse, error) {
	_, err := customerService.customerRepository.GetByID(ctx, customerRequest.ID)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	customerModel, err := customerService.customerRepository.Update(ctx, customerRequest)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	return customerModel.Entity2Response(), nil
}

func (customerService *CustomerService) CreateCustomer(ctx context.Context, customer models.CustomerRequest) (*models.CustomerResponse, error) {
	_, err := customerService.customerRepository.GetByEmail(ctx, customer.Email)
	if err == nil {
		return nil, errors.Errorf("Email %s Exist", customer.Email)
	}

	customerModel, err := customerService.customerRepository.Save(ctx, customer)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return customerModel.Entity2Response(), nil
}

func (customerService *CustomerService) GetCertByID(ctx context.Context, id int) (*models.CustomerResponse, error) {
	return nil, nil
}

func (customerService *CustomerService) GetAllByClinicID(ctx context.Context, id string) ([]*models.CustomerResponse, error) {
	customers, err := customerService.customerRepository.GetAllByClinicID(ctx, id)
	if err != nil {
		return nil, err
	}

	result := make([]*models.CustomerResponse, 0)
	for _, v := range customers {
		result = append(result, v.Entity2Response())
	}

	return result, nil
}
