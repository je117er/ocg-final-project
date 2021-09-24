package booking

import (
	"context"
	"github.com/je117er/ocg-final-project/internal/models"
)

type Repository interface {
	GetByCustomerID(ctx context.Context, customerID int) (*models.Booking, error)
	GetByID(ctx context.Context, id int) (*models.Booking, error)
	UpdateCompleted(ctx context.Context, id int, completed int) error
}
