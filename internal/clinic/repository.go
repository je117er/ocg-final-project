package clinic

import (
	"context"
	"github.com/je117er/ocg-final-project/internal/models"
)

type Repository interface {
	Fetch(ctx context.Context) ([]*models.Clinic, error)
	GetByID(ctx context.Context, id string) (*models.Clinic, error)
	Update(ctx context.Context, clinic models.ClinicRequest) (*models.Clinic, error)
	GetClinicByVaccine(ctx context.Context, vaccine string) ([]*models.ClinicByVaccine, error)
}
