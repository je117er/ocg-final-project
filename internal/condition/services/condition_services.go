package constraint

import (
	"context"
	"github.com/je117er/ocg-final-project/internal/condition"
	"github.com/je117er/ocg-final-project/internal/models"
	"github.com/je117er/ocg-final-project/internal/utils"
)

type ConditionService struct {
	conditionRepo condition.Repository
}

func NewConditionService(conditionRepo condition.Repository) condition.Service {
	return &ConditionService{
		conditionRepo: conditionRepo,
	}
}

var logger = utils.SugarLog()

func (service *ConditionService) ByCustomerID(ctx context.Context, id int) ([]*models.MedicalConditionResponse, error) {
	res, err := service.conditionRepo.ByCustomerID(ctx, id)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	result := make([]*models.MedicalConditionResponse, 0)
	for _, v := range res {
		result = append(result, v.ModelToResponse())
	}
	return result, nil
}
