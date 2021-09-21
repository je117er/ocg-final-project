package clinic

import (
	"context"
	"github.com/je117er/ocg-final-project/internal/models"
)

type Repository interface {
	Fetch(ctx context.Context) ([]*models.Clinic, error)
	/*GetByID(ctx context.Context, id int) (*models.Clinic, error)
	Update(ctx context.Context, clinic models.Clinic) (*models.Clinic, error)*/
}