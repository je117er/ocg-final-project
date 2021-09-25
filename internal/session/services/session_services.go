package session

import (
	"context"
	"github.com/je117er/ocg-final-project/internal/models"
	"github.com/je117er/ocg-final-project/internal/session"
	"github.com/je117er/ocg-final-project/internal/utils"
)

type SessionService struct {
	SessionRepository session.Repository
}

func NewSessionService(sessionRepository session.Repository) session.Service {
	return &SessionService{
		SessionRepository: sessionRepository,
	}
}

var logger = utils.SugarLog()

func (service *SessionService) GetSessionByClinic(ctx context.Context, clinicName string) ([]*models.SessionByClinicResponse, error) {
	results, err := service.SessionRepository.GetSessionByClinic(ctx, clinicName)
	if err != nil {
		return nil, err
	}

	responses := make([]*models.SessionByClinicResponse, 0)
	for _, res := range results {
		responses = append(responses, res.ModelToResponse())
	}
	return responses, err
}
func (service *SessionService) GetAllSessionInMonth(ctx context.Context, month int) ([]*models.SessionCapacityResponse, error) {
	result, err := service.SessionRepository.GetAllSessionInMonth(ctx, month)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	response := make([]*models.SessionCapacityResponse, 0)
	for _, v := range result {
		response = append(response, v.Entity2Response())
	}
	return response, nil
}

func (service *SessionService) UpdateSession(ctx context.Context, requests []*models.SessionCapacityRequest) error {
	for _, v := range requests {
		service.SessionRepository.Update(ctx, v)
	}
	return nil
}
