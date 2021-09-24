package server

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	adminController "github.com/je117er/ocg-final-project/internal/admin/controller"
	adminRepository "github.com/je117er/ocg-final-project/internal/admin/repository"
	adminService "github.com/je117er/ocg-final-project/internal/admin/services"
	bookingController "github.com/je117er/ocg-final-project/internal/booking/controller"
	bookingRepository "github.com/je117er/ocg-final-project/internal/booking/repository"
	bookingService "github.com/je117er/ocg-final-project/internal/booking/services"
	clinicController "github.com/je117er/ocg-final-project/internal/clinic/controller"
	clinicRepository "github.com/je117er/ocg-final-project/internal/clinic/repository"
	clinicService "github.com/je117er/ocg-final-project/internal/clinic/services"
	conditionController "github.com/je117er/ocg-final-project/internal/condition/controller"
	conditionRepository "github.com/je117er/ocg-final-project/internal/condition/repository"
	conditionService "github.com/je117er/ocg-final-project/internal/condition/services"
	constraintController "github.com/je117er/ocg-final-project/internal/constraint/controller"
	constraintRepositoty "github.com/je117er/ocg-final-project/internal/constraint/repository"
	constraintService "github.com/je117er/ocg-final-project/internal/constraint/services"
	customerController "github.com/je117er/ocg-final-project/internal/customer/controller"
	customerRepository "github.com/je117er/ocg-final-project/internal/customer/repository"
	customerService "github.com/je117er/ocg-final-project/internal/customer/services"
	"github.com/je117er/ocg-final-project/internal/middleware"
	productController "github.com/je117er/ocg-final-project/internal/product/controllers"
	productRepository "github.com/je117er/ocg-final-project/internal/product/repository"
	productService "github.com/je117er/ocg-final-project/internal/product/services"
	sessionController "github.com/je117er/ocg-final-project/internal/session/controller"
	sessionRepository "github.com/je117er/ocg-final-project/internal/session/repository"
	sessionService "github.com/je117er/ocg-final-project/internal/session/services"
	stockRepository "github.com/je117er/ocg-final-project/internal/stock/repository"
	"github.com/je117er/ocg-final-project/internal/utils"
	"github.com/rs/cors"
	"net/http"
)

var logger = utils.SugarLog()

func InitServer() error {
	logger.Info("Start server ...")
	DB, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"admin", "123456Aa@", "localhost", "6033", "vaccine-covid-19"))

	if err != nil {
		logger.Fatal(err)
	}
	defer DB.Close()

	ctx := context.Background()

	productRepo := productRepository.NewProductRepository(DB)
	customerRepo := customerRepository.NewCustomerRepository(DB)
	constraintRepo := constraintRepositoty.NewConstraintRepository(DB)
	clinicRepo := clinicRepository.NewClinicRepository(DB)
	adminRepo := adminRepository.NewAdminRepository(DB)
	conditionRepo := conditionRepository.NewConditionRepository(DB)
	bookingRepo := bookingRepository.NewBookingRepository(DB)
	stockRepo := stockRepository.NewStockRepository(DB)
	sessionRepo := sessionRepository.NewSessionRepository(DB)

	productServ := productService.NewProductService(productRepo)
	customerServ := customerService.NewCustomerService(customerRepo)
	constraintServ := constraintService.NewConstraintService(constraintRepo)
	clinicServ := clinicService.NewClinicService(clinicRepo)
	adminServ := adminService.NewAdminService(adminRepo)
	conditionServ := conditionService.NewConditionService(conditionRepo)
	bookingServ := bookingService.NewBookingService(bookingRepo, stockRepo)
	sessionServ := sessionService.NewSessionService(sessionRepo)

	r := mux.NewRouter()
	productController.NewProductController(productServ, ctx, r)
	customerController.NewCustomerController(customerServ, ctx, r)
	constraintController.NewConstraintController(constraintServ, ctx, r)
	clinicController.NewClinicController(clinicServ, ctx, r)
	adminController.NewAdminController(adminServ, ctx, r)
	conditionController.NewConditionController(conditionServ, ctx, r)
	bookingController.NewBookingController(bookingServ, ctx, r)
	sessionController.NewAdminCustomerController(sessionServ, ctx, r)

	customerController.NewAdminCustomerController(customerServ, ctx, r)
	s := r.PathPrefix("/admin").Subrouter()
	s.Use(middleware.JWTVerify)

	// cors.Default() sets up the middleware with default options being
	// all origins accepted with simple methods (GET, POST)
	// references: https://github.com/rs/cors
	handler := cors.AllowAll().Handler(r)
	err = http.ListenAndServe(":8088", handler)
	if err != nil {
		return err
	}
	return nil
}
