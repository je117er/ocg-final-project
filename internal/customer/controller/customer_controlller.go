package controller

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/je117er/ocg-final-project/internal/customer"
	"github.com/je117er/ocg-final-project/internal/models"
	"github.com/je117er/ocg-final-project/internal/utils"
	"net/http"
)

type CustomerController struct {
	CustomerService customer.Service
	ctx             context.Context
}

var logger = utils.SugarLog()

func NewCustomerController(service customer.Service, ctx context.Context, router *mux.Router) {
	controller := CustomerController{service, ctx}
	router.Methods(http.MethodGet).Path("/customer").Queries("email", "{email}").HandlerFunc(controller.GetCustomerByEmail)
	router.Methods(http.MethodPut).Path("/customer").HandlerFunc(controller.UpdateCustomer)
	router.Methods(http.MethodPost).Path("/customer").HandlerFunc(controller.CreateCustomer)
	router.Methods(http.MethodGet).Path("/customer").PathPrefix("/{id}/cert").HandlerFunc(controller.GetCertByID)
}

func (controller *CustomerController) GetCustomerByEmail(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")

	c, err := controller.CustomerService.GetByEmail(controller.ctx, email)
	if err != nil {
		logger.Error(err)
		utils.ResponseWithJson(w, http.StatusInternalServerError, models.ResponseError{Message: err.Error()})

		return
	}

	utils.ResponseWithJson(w, http.StatusOK, c)
}

func (controller *CustomerController) UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	var customerRequest models.CustomerRequest
	if err := json.NewDecoder(r.Body).Decode(&customerRequest); err != nil {
		utils.ResponseWithJson(w, http.StatusBadRequest, models.ResponseError{Message: "Invalid body"})
		return
	}

	customerResponse, err := controller.CustomerService.UpdateCustomer(controller.ctx, customerRequest)
	if err != nil {
		logger.Error(err)
		utils.ResponseWithJson(w, http.StatusInternalServerError, models.ResponseError{Message: "Internal Server Error"})

		return
	}

	utils.ResponseWithJson(w, http.StatusOK, customerResponse)
}

func (controller *CustomerController) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var customerRequest models.CustomerRequest
	if err := json.NewDecoder(r.Body).Decode(&customerRequest); err != nil {
		utils.ResponseWithJson(w, http.StatusBadRequest, models.ResponseError{Message: "Invalid body"})
		return
	}

	customerResponse, err := controller.CustomerService.CreateCustomer(controller.ctx, customerRequest)
	if err != nil {
		utils.ResponseWithJson(w, http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
		return
	}

	utils.ResponseWithJson(w, http.StatusOK, customerResponse)
}

func (controller *CustomerController) GetCertByID(w http.ResponseWriter, r *http.Request) {

}
