package clinic

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/je117er/ocg-final-project/internal/booking"
	"github.com/je117er/ocg-final-project/internal/models"
	"github.com/je117er/ocg-final-project/internal/utils"
	"net/http"
	"strconv"
)

type BookingController struct {
	BookingService booking.Service
	ctx            context.Context
}

var logger = utils.SugarLog()

func NewBookingController(service booking.Service, ctx context.Context, router *mux.Router) {
	controller := BookingController{service, ctx}
	router.Methods(http.MethodGet).Path("/admin/booking").Queries("customerID", "{customerID}").HandlerFunc(controller.GetVaccinationByCustomerID)
	router.Methods(http.MethodPut).Path("/admin/booking").HandlerFunc(controller.UpdateDoesVaccination)
}

func (controller *BookingController) GetVaccinationByCustomerID(w http.ResponseWriter, r *http.Request) {
	customerID, err := strconv.Atoi(r.URL.Query().Get("customerID"))
	if err != nil {
		logger.Error(err)
		utils.ResponseWithJson(w, http.StatusInternalServerError, models.ResponseError{Message: err.Error()})

		return
	}

	response, err := controller.BookingService.GetVaccinationByCustomerID(controller.ctx, customerID)
	if err != nil {
		logger.Error(err)
		utils.ResponseWithJson(w, http.StatusInternalServerError, models.ResponseError{Message: err.Error()})

		return
	}

	utils.ResponseWithJson(w, http.StatusOK, response)
}

func (controller *BookingController) UpdateDoesVaccination(w http.ResponseWriter, r *http.Request) {
	var vaccinationRequest models.VaccinationRequest
	if err := json.NewDecoder(r.Body).Decode(&vaccinationRequest); err != nil {
		logger.Error(err)
		utils.ResponseWithJson(w, http.StatusBadRequest, models.ResponseError{Message: "Invalid body"})
		return
	}

	response, err := controller.BookingService.UpdateCustomer(controller.ctx, vaccinationRequest)
	if err != nil {
		logger.Error(err)
		utils.ResponseWithJson(w, http.StatusInternalServerError, models.ResponseError{Message: err.Error()})

		return
	}

	utils.ResponseWithJson(w, http.StatusOK, response)
}
