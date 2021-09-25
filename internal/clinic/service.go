package clinic

import (
	"context"
	"github.com/je117er/ocg-final-project/internal/models"
)

type Service interface {
	GetAll(ctx context.Context) ([]*models.ClinicResponse, error)
	FindByID(ctx context.Context, id string) (*models.ClinicResponse, error)
	UpdateClinic(ctx context.Context, customerRequest models.ClinicRequest) (*models.ClinicResponse, error)
	GetClinicByVaccine(ctx context.Context, vaccine string) ([]*models.ClinicByVaccineResponse, error)
}
