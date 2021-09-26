package booking

import (
	"context"
	"database/sql"
	"github.com/je117er/ocg-final-project/internal/models"
)

type Repository interface {
	GetByCustomerID(ctx context.Context, customerID int) (*models.Booking, error)
	GetByID(ctx context.Context, id int) (*models.Booking, error)
	UpdateCompleted(ctx context.Context, id int, completed int) error
	SaveOrder(ctx context.Context, r models.OrderRequest, customerID int) (*sql.Tx, error)
	UpdateIsSendRemindEmail(ctx context.Context, id int) error
}
