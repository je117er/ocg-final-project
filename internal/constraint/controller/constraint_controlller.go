package constraint

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/je117er/ocg-final-project/internal/constraint"
	"github.com/je117er/ocg-final-project/internal/models"
	"github.com/je117er/ocg-final-project/internal/utils"
	"net/http"
)

type ConstraintController struct {
	ConstraintService constraint.Service
	ctx               context.Context
}

var logger = utils.SugarLog()

func NewConstraintController(service constraint.Service, ctx context.Context, router *mux.Router) {
	controller := ConstraintController{service, ctx}
	router.Methods(http.MethodGet).Path("/constraints").HandlerFunc(controller.GetAll)
}

func (controller *ConstraintController) GetAll(w http.ResponseWriter, r *http.Request) {
	response, err := controller.ConstraintService.GetAll(controller.ctx)
	if err != nil {
		logger.Error(err)
		utils.ResponseWithJson(w, http.StatusInternalServerError, models.ResponseError{Message: err.Error()})

		return
	}

	utils.ResponseWithJson(w, http.StatusOK, response)
}
