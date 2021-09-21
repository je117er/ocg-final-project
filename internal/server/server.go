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
	customerController "github.com/je117er/ocg-final-project/internal/customer/controller"
	customerRepository "github.com/je117er/ocg-final-project/internal/customer/repository"
	customerService "github.com/je117er/ocg-final-project/internal/customer/services"
	"github.com/je117er/ocg-final-project/internal/middleware"
	productController "github.com/je117er/ocg-final-project/internal/product/controllers"
	productRepository "github.com/je117er/ocg-final-project/internal/product/repository"
	productService "github.com/je117er/ocg-final-project/internal/product/services"
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
	adminRepo := adminRepository.NewAdminRepository(DB)

	productServ := productService.NewProductService(productRepo)
	customerServ := customerService.NewCustomerService(customerRepo)
	adminServ := adminService.NewAdminService(adminRepo)

	r := mux.NewRouter()
	adminController.NewAdminController(adminServ, ctx, r)
	customerController.NewCustomerController(customerServ, ctx, r)
	productController.NewProductController(productServ, ctx, r)
	s := r.PathPrefix("/admin").Subrouter()
	s.Use(middleware.JWTVerify)
	//productController.NewProductController(productServ, ctx, s)
	//customerController.NewCustomerController(customerServ, ctx, r)

	// cors.Default() sets up the middleware with default options being
	// all origins accepted with simple methods (GET, POST)
	// references: https://github.com/rs/cors
	handler := cors.Default().Handler(r)
	err = http.ListenAndServe(":8088", handler)
	if err != nil {
		return err
	}
	return nil
}
