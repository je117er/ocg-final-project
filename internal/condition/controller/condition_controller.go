package constraint

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/je117er/ocg-final-project/internal/condition"
	"github.com/je117er/ocg-final-project/internal/utils"
	"net/http"
	"strconv"
)

type ConditionController struct {
	ConditionService condition.Service
	ctx              context.Context
}

var logger = utils.SugarLog()

func NewConditionController(service condition.Service, ctx context.Context, router *mux.Router) {
	controller := ConditionController{service, ctx}
	router.Methods(http.MethodGet).Path("/admin/medicalHistory").Queries("customerID", "{customerID}").HandlerFunc(controller.GetMedicalHistory)
}

func (controller *ConditionController) GetMedicalHistory(w http.ResponseWriter, r *http.Request) {
	customerID, err := strconv.Atoi(r.URL.Query().Get("customerID"))
	if err != nil {
		logger.Error(err)
		utils.ResponseWithJson(w, http.StatusBadRequest, utils.ResponseError{Message: err.Error()})
		return
	}
	conditions, err := controller.ConditionService.ByCustomerID(controller.ctx, customerID)
	if err != nil {
		logger.Error(err)
		utils.ResponseWithJson(w, http.StatusNotFound, utils.ResponseError{Message: err.Error()})
		return
	}
	utils.ResponseWithJson(w, http.StatusOK, conditions)
}
