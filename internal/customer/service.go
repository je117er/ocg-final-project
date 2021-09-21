package customer

import (
	"context"
	"github.com/je117er/ocg-final-project/internal/models"
)

type Service interface {
	GetByEmail(ctx context.Context, email string) (*models.CustomerResponse, error)
}
