package condition

import (
	"context"
	"github.com/je117er/ocg-final-project/internal/models"
)

type Service interface {
	ByCustomerID(ctx context.Context, id int) ([]*models.MedicalConditionResponse, error)
}
