package customer

import (
	"context"
	"github.com/je117er/ocg-final-project/internal/models"
)

type Repository interface {
	GetByEmail(ctx context.Context, email string) (*models.Customer, error)
}
