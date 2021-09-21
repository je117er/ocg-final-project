package customer

import (
	"context"
	"github.com/je117er/ocg-final-project/internal/models"
)

type Service interface {
	GetByEmail(ctx context.Context, email string) (*models.CustomerResponse, error)
	UpdateCustomer(ctx context.Context, customer models.CustomerRequest) (*models.CustomerResponse, error)
	CreateCustomer(ctx context.Context, customer models.CustomerRequest) (*models.CustomerResponse, error)
	GetCertByID(ctx context.Context, id int) (*models.CustomerResponse, error)
	GetAllByClinicID(ctx context.Context, id string) ([]*models.CustomerResponse, error)
	GetByID(ctx context.Context, id int) (*models.CustomerResponse, error)
	GetAll(ctx context.Context) ([]*models.CustomerResponse, error)
}
