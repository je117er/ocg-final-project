package session

import (
	"context"
	"github.com/je117er/ocg-final-project/internal/models"
)

type Repository interface {
	GetAllSessionInMonth(ctx context.Context, month int) ([]*models.SessionCapacity, error)
	Update(ctx context.Context, request *models.SessionCapacityRequest) error
	GetSessionByClinic(ctx context.Context, clinicName string) ([]*models.SessionByClinic, error)
}
