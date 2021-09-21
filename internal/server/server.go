package server

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	clinicController "github.com/je117er/ocg-final-project/internal/clinic/controller"
	clinicRepository "github.com/je117er/ocg-final-project/internal/clinic/repository"
	clinicService "github.com/je117er/ocg-final-project/internal/clinic/services"
	constraintController "github.com/je117er/ocg-final-project/internal/constraint/controller"
	constraintRepositoty "github.com/je117er/ocg-final-project/internal/constraint/repository"
	constraintService "github.com/je117er/ocg-final-project/internal/constraint/services"
	customerController "github.com/je117er/ocg-final-project/internal/customer/controller"
	customerRepository "github.com/je117er/ocg-final-project/internal/customer/repository"
	customerService "github.com/je117er/ocg-final-project/internal/customer/services"
	productController "github.com/je117er/ocg-final-project/internal/product/controllers"
	productRepository "github.com/je117er/ocg-final-project/internal/product/repository"
	productService "github.com/je117er/ocg-final-project/internal/product/services"
	"github.com/je117er/ocg-final-project/internal/utils"
	"github.com/rs/cors"
	"log"

	"net/http"
)

var logger = utils.SugarLog()

func Middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("middleware", r.URL)
		h.ServeHTTP(w, r)
	})
}

func InitServer() error {
	logger.Info("Start server ...")
	DB, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"root", "123456Aa@", "localhost", "6033", "vaccine-covid-19"))
	if err != nil {
		logger.Fatal(err)
	}
	defer DB.Close()

	/*ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()*/
	ctx := context.Background()

	productRepo := productRepository.NewProductRepository(DB)
	customerRepo := customerRepository.NewCustomerRepository(DB)
	constraintRepo := constraintRepositoty.NewConstraintRepository(DB)
	clinicRepo := clinicRepository.NewClinicRepository(DB)

	productService := productService.NewProductService(productRepo)
	customerService := customerService.NewCustomerService(customerRepo)
	constraintService := constraintService.NewConstraintService(constraintRepo)
	clinicService := clinicService.NewClinicService(clinicRepo)

	r := mux.NewRouter()
	r.Use(Middleware)
	productController.NewProductController(productService, ctx, r)
	customerController.NewCustomerController(customerService, ctx, r)
	constraintController.NewConstraintController(constraintService, ctx, r)
	clinicController.NewClinicController(clinicService, ctx, r)

	// cors.Default() setup the middleware with default options being
	// all origins accepted with simple methods (GET, POST)
	// references: https://github.com/rs/cors
	handler := cors.Default().Handler(r)
	err = http.ListenAndServe(":8088", handler)
	if err != nil {
		return err
	}
	return nil
}
