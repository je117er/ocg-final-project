package constraint

import (
	"context"
	"database/sql"
	"github.com/je117er/ocg-final-project/internal/models"
)

type Repository interface {
	Fetch(ctx context.Context) ([]*models.Constraint, error)
	SaveConstraint(ctx context.Context, conditions []models.ConditionOrderRequest, customerID int) (*sql.Tx, error)
}
