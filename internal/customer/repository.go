package customer

import (
	"context"
	"github.com/je117er/ocg-final-project/internal/models"
)

type Repository interface {
	GetByEmail(ctx context.Context, email string) (*models.Customer, error)
	GetByID(ctx context.Context, id int) (*models.Customer, error)
	Save(ctx context.Context, customer models.CustomerRequest) (*models.Customer, error)
	Update(ctx context.Context, customer models.CustomerRequest) (*models.Customer, error)
	GetCertByID(ctx context.Context, id int) (*models.CustomerResponse, error)
	GetAllByClinicID(ctx context.Context, id string) ([]*models.Customer, error)
}
