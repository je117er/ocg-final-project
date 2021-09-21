package controller

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/je117er/ocg-final-project/internal/customer"
	"github.com/je117er/ocg-final-project/internal/utils"
	"net/http"
)

type ResponseError struct {
	Message string `json:"message"`
}

type CustomerController struct {
	CustomerService customer.Service
	ctx             context.Context
}

var logger = utils.SugarLog()

func NewCustomerController(service customer.Service, ctx context.Context, router *mux.Router) {
	controller := CustomerController{service, ctx}
	router.Methods(http.MethodGet).Path("/customer").Queries("email", "{email}").HandlerFunc(controller.GetCustomerByEmail)
}

func (controller *CustomerController) GetCustomerByEmail(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")

	c, err := controller.CustomerService.GetByEmail(controller.ctx, email)
	if err != nil {
		logger.Error(err)
		utils.ResponseWithJson(w, http.StatusInternalServerError, ResponseError{"Internal Server Error"})

		return
	}

	utils.ResponseWithJson(w, http.StatusOK, c)
}
