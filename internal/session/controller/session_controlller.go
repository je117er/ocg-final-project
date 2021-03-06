package session

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/je117er/ocg-final-project/internal/models"
	"github.com/je117er/ocg-final-project/internal/session"
	"github.com/je117er/ocg-final-project/internal/utils"
	"net/http"
	"strconv"
)

type SessionController struct {
	SessionService session.Service
	ctx            context.Context
}

var logger = utils.SugarLog()

func NewAdminSessionController(service session.Service, ctx context.Context, router *mux.Router) {
	controller := SessionController{service, ctx}
	router.Methods(http.MethodGet).Path("/session").Queries("month", "{month}").HandlerFunc(controller.GetAllSessionInMonth)
	router.Methods(http.MethodPut).Path("/session").HandlerFunc(controller.UpdateSession)
}
func NewSessionController(service session.Service, ctx context.Context, router *mux.Router) {
	controller := SessionController{service, ctx}
	router.Methods(http.MethodGet).Path("/session").Queries("clinicName", "{clinicName}").HandlerFunc(controller.GetSessionByClinic)
}
func (controller *SessionController) GetSessionByClinic(w http.ResponseWriter, r *http.Request) {
	clinicName := mux.Vars(r)["clinicName"]
	results, err := controller.SessionService.GetSessionByClinic(controller.ctx, clinicName)
	if err != nil {
		logger.Error(err)
		utils.ResponseWithJson(w, http.StatusBadRequest, utils.ResponseError{Message: err.Error()})
		return
	}
	utils.ResponseWithJson(w, http.StatusOK, results)
}

func (controller *SessionController) GetAllSessionInMonth(w http.ResponseWriter, r *http.Request) {
	month, err := strconv.Atoi(r.URL.Query().Get("month"))
	if err != nil {
		logger.Error(err)
		utils.ResponseWithJson(w, http.StatusInternalServerError, "Invalid Param")

		return
	}

	res, err := controller.SessionService.GetAllSessionInMonth(controller.ctx, month)
	if err != nil {
		logger.Error(err)
		utils.ResponseWithJson(w, http.StatusInternalServerError, models.ResponseError{Message: err.Error()})

		return
	}

	utils.ResponseWithJson(w, http.StatusOK, res)
}

func (controller *SessionController) UpdateSession(w http.ResponseWriter, r *http.Request) {
	var sessionRequests []*models.SessionCapacityRequest
	err := json.NewDecoder(r.Body).Decode(&sessionRequests)
	if err != nil {
		logger.Error(err)
		utils.ResponseWithJson(w, http.StatusInternalServerError, "Invalid Body")

		return
	}

	err = controller.SessionService.UpdateSession(controller.ctx, sessionRequests)
	if err != nil {
		logger.Error(err)
		utils.ResponseWithJson(w, http.StatusInternalServerError, models.ResponseError{Message: err.Error()})

		return
	}

	utils.ResponseWithJson(w, http.StatusOK, nil)

}
