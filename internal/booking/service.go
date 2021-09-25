package booking

import (
	"context"
	"github.com/je117er/ocg-final-project/internal/models"
)

type Service interface {
	GetVaccinationByCustomerID(ctx context.Context, customerID int) ([]models.VaccinationResponse, error)
	UpdateCustomer(ctx context.Context, customer models.VaccinationRequest) ([]models.VaccinationResponse, error)
	InsertOrder(ctx context.Context, r models.OrderRequest) error
}
