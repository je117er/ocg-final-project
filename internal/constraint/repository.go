package constraint

import (
	"context"
	"github.com/je117er/ocg-final-project/internal/models"
)

type Repository interface {
	Fetch(ctx context.Context) ([]*models.Constraint, error)
}
