package condition

import (
	"context"
	"github.com/je117er/ocg-final-project/internal/models"
)

type Repository interface {
	ByCustomerID(ctx context.Context, id int) ([]*models.MedicalCondition, error)
}
