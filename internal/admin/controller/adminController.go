package controller

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/je117er/ocg-final-project/internal/admin"
	"github.com/je117er/ocg-final-project/internal/models"
	"github.com/je117er/ocg-final-project/internal/utils"
	"net/http"
)

type ResponseError struct {
	Message string `json:"message"`
}

type AdminController struct {
	AdminService admin.Service
	ctx          context.Context
}

var logger = utils.SugarLog()

func NewAdminController(as admin.Service, ctx context.Context, r *mux.Router) {
	controller := &AdminController{as, ctx}
	r.Methods("POST").Path("/login").HandlerFunc(controller.Login)
}

func (as *AdminController) Login(w http.ResponseWriter, r *http.Request) {
	user := &models.LoginRequest{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.ResponseWithJson(w, http.StatusBadRequest, ResponseError{"Invalid Request"})
		logger.Error(err)
		return
	}
	resp := as.AdminService.Authenticate(as.ctx, user.Username, user.Password)

	err := resp.Message
	switch err {
	case admin.ErrInvalidUsername:
		utils.ResponseWithJson(w, http.StatusBadRequest, admin.ErrInvalidUsername)
		return
	case admin.ErrInvalidPassword:
		utils.ResponseWithJson(w, http.StatusBadRequest, admin.ErrInvalidPassword)
		return
	case admin.ErrGeneratingToken:
		utils.ResponseWithJson(w, http.StatusInternalServerError, admin.ErrGeneratingToken)
		return
	}
	utils.ResponseWithJson(w, http.StatusOK, resp)
}
