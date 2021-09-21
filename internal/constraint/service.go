package constraint

import (
	"context"
	"github.com/je117er/ocg-final-project/internal/models"
)

type Service interface {
	GetAll(ctx context.Context) ([]*models.ConstraintResponse, error)
}
