package controller

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/je117er/ocg-final-project/internal/customer"
	"github.com/je117er/ocg-final-project/internal/models"
	"github.com/je117er/ocg-final-project/internal/utils"
	"net/http"
	"strconv"
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

func NewAdminCustomerController(service customer.Service, ctx context.Context, router *mux.Router) {
	controller := CustomerController{service, ctx}
	router.Methods(http.MethodGet).Path("/admin/customers").HandlerFunc(controller.GetCustomers)
	router.Methods(http.MethodGet).Path("/admin/customer").Queries("clinicId", "{clinicId}").HandlerFunc(controller.GetCustomerByClinicId)
	router.Methods(http.MethodPut).Path("/admin/customer").Queries("clinicId", "{clinicId}").HandlerFunc(controller.UpdateCustomer)
	router.Methods(http.MethodGet).Path("/admin/customer/{id:[0-9]+}").HandlerFunc(controller.GetCustomerByID)
	router.Methods(http.MethodPut).Path("/admin/customer").HandlerFunc(controller.UpdateCustomer)

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

func (controller *CustomerController) GetCustomers(w http.ResponseWriter, r *http.Request) {
	c, err := controller.CustomerService.GetAll(controller.ctx)
	if err != nil {
		logger.Error(err)
		utils.ResponseWithJson(w, http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
		return
	}
	utils.ResponseWithJson(w, http.StatusOK, c)
}

func (controller *CustomerController) GetCustomerByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		logger.Error(err)
		utils.ResponseWithJson(w, http.StatusBadRequest, models.ResponseError{Message: err.Error()})
		return
	}

	cust, err := controller.CustomerService.GetByID(controller.ctx, id)
	if err != nil {
		utils.ResponseWithJson(w, http.StatusNotFound, models.ResponseError{Message: err.Error()})
		return
	}
	utils.ResponseWithJson(w, http.StatusOK, cust)
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

func (controller *CustomerController) GetCustomerByClinicId(w http.ResponseWriter, r *http.Request) {
	clinicId := r.URL.Query().Get("clinicId")

	c, err := controller.CustomerService.GetAllByClinicID(controller.ctx, clinicId)
	if err != nil {
		logger.Error(err)
		utils.ResponseWithJson(w, http.StatusInternalServerError, models.ResponseError{Message: err.Error()})

		return
	}

	utils.ResponseWithJson(w, http.StatusOK, c)
}
