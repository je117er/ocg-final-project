package clinic

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/je117er/ocg-final-project/internal/booking"
	"github.com/je117er/ocg-final-project/internal/models"
	"github.com/je117er/ocg-final-project/internal/utils"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/checkout/session"
	"log"
	"net/http"
	"strconv"
)

type BookingController struct {
	BookingService booking.Service
	ctx            context.Context
}

var (
	logger = utils.SugarLog()
)

type SuccessOrderMessage struct {
	Message string `json:"message"`
}

func NewAdminBookingController(service booking.Service, ctx context.Context, router *mux.Router) {
	controller := BookingController{service, ctx}
	router.Methods(http.MethodGet).Path("/booking").Queries("customerID", "{customerID}").HandlerFunc(controller.GetVaccinationByCustomerID)
	router.Methods(http.MethodPut).Path("/booking").HandlerFunc(controller.UpdateDoesVaccination)
}

func NewBookingController(service booking.Service, ctx context.Context, router *mux.Router) {
	controller := BookingController{service, ctx}
	router.Methods(http.MethodPost).Path("/order").HandlerFunc(controller.PostOrder)
	router.Methods(http.MethodPost).Path("/create-checkout-session").HandlerFunc(controller.CreateCheckoutSession)
	router.Methods(http.MethodGet).Path("/booking").Queries("customerID", "{customerID}").HandlerFunc(controller.GetVaccinationByCustomerID)
	router.Methods(http.MethodPut).Path("/booking").HandlerFunc(controller.UpdateDoesVaccination)
}

func (controller BookingController) CreateCheckoutSession(w http.ResponseWriter, r *http.Request) {
	// secret key
	stripe.Key = "sk_test_51JRxRHILBKw97Kynl4p3dkVtCeuoaPmJHcuXOqr54StCQcwxGlGxQkvSFScVxPqE1B03NnxYDj1FlN3ImP0oX4jI00kvpGyjWV"
	var req models.CreateCheckoutSessionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Error(err)
		utils.ResponseWithJson(w, http.StatusBadRequest, utils.ResponseError{Message: err.Error()})
		return
	}
	params := &stripe.CheckoutSessionParams{
		Customer: stripe.String(string(rune(req.CustomerID))),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				Price:    stripe.String(string(rune(req.TotalBill))),
				Quantity: stripe.Int64(int64(req.Quantity)),
			},
		},
		PaymentMethodTypes: stripe.StringSlice([]string{
			"card",
		}),
		Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL: stripe.String(req.SuccessURL),
		CancelURL:  stripe.String(req.CancelURL),
	}
	s, err := session.New(params)
	if err != nil {
		log.Printf("session.New: %v", err)
	}
	http.Redirect(w, r, s.URL, http.StatusSeeOther)
}

func (controller BookingController) PostOrder(w http.ResponseWriter, r *http.Request) {
	var request models.OrderRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		logger.Error(err)
		utils.ResponseWithJson(w, http.StatusBadRequest, utils.ResponseError{Message: err.Error()})
		return
	}

	if err := controller.BookingService.InsertOrder(controller.ctx, request); err != nil {
		logger.Error(err)
		utils.ResponseWithJson(w, http.StatusBadRequest, utils.ResponseError{Message: err.Error()})
		return
	}

	utils.ResponseWithJson(w, http.StatusAccepted, SuccessOrderMessage{Message: "Success placing an order!"})
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
