package clinic

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/je117er/ocg-final-project/internal/clinic"
	"github.com/je117er/ocg-final-project/internal/models"
	"github.com/je117er/ocg-final-project/internal/utils"
	"net/http"
)

type ClinicController struct {
	ClinicService clinic.Service
	ctx           context.Context
}

var logger = utils.SugarLog()

func NewClinicController(service clinic.Service, ctx context.Context, router *mux.Router) {
	controller := ClinicController{service, ctx}
	router.Methods(http.MethodGet).Path("/admin/clinics").HandlerFunc(controller.GetAll)
}

func (controller *ClinicController) GetAll(w http.ResponseWriter, r *http.Request) {
	response, err := controller.ClinicService.GetAll(controller.ctx)
	if err != nil {
		logger.Error(err)
		utils.ResponseWithJson(w, http.StatusInternalServerError, models.ResponseError{Message: err.Error()})

		return
	}

	utils.ResponseWithJson(w, http.StatusOK, response)
}
