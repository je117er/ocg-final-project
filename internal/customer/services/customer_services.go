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
