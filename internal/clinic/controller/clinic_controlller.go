package clinic

import (
	"context"
	"encoding/json"
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
	router.Methods(http.MethodGet).Path("/clinics").HandlerFunc(controller.GetAll)
	router.Methods(http.MethodGet).Path("/clinic/{id}/info").HandlerFunc(controller.GetByClinicID)
	router.Methods(http.MethodPut).Path("/clinic/{id}/info").HandlerFunc(controller.UpdateClinic)
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

func (controller *ClinicController) GetByClinicID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	response, err := controller.ClinicService.FindByID(controller.ctx, id)
	if err != nil {
		logger.Error(err)
		utils.ResponseWithJson(w, http.StatusInternalServerError, models.ResponseError{Message: err.Error()})

		return
	}

	utils.ResponseWithJson(w, http.StatusOK, response)
}

func (controller *ClinicController) UpdateClinic(w http.ResponseWriter, r *http.Request) {
	var clinicRequest models.ClinicRequest
	if err := json.NewDecoder(r.Body).Decode(&clinicRequest); err != nil {
		utils.ResponseWithJson(w, http.StatusBadRequest, models.ResponseError{Message: err.Error()})
		return
	}

	clinicResponse, err := controller.ClinicService.UpdateClinic(controller.ctx, clinicRequest)
	if err != nil {
		logger.Error(err)
		utils.ResponseWithJson(w, http.StatusInternalServerError, models.ResponseError{Message: "Internal Server Error"})

		return
	}

	utils.ResponseWithJson(w, http.StatusOK, clinicResponse)

}
