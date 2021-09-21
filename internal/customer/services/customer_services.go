package services

import (
	"context"
	"github.com/je117er/ocg-final-project/internal/customer"
	"github.com/je117er/ocg-final-project/internal/models"
)

type CustomerService struct {
	customerRepository customer.Repository
}

func NewCustomerService(customerRepository customer.Repository) customer.Service {
	return &CustomerService{
		customerRepository: customerRepository,
	}
}

func (customerService *CustomerService) GetByEmail(ctx context.Context, email string) (*models.CustomerResponse, error) {
	res, err := customerService.customerRepository.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return res.Entity2Response(), nil
}
