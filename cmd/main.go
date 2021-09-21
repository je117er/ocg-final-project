package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/je117er/ocg-final-project/internal/product/repository"
	"github.com/je117er/ocg-final-project/internal/utils"
	"log"
	"time"
)

func main() {
	sugarLogger := utils.SugarLog()
	sugarLogger.Info("Test log info")
	sugarLogger.Debug("Test log debug")
	sugarLogger.Error("Test log err")

	DB, err := sql.Open("mysql", "root:123456Aa@@tcp(localhost:6033)/vaccine-covid-19?charset=utf8mb4&parseTime=True&loc=Local")
	//DB, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", "root", "123456Aa@", "localhost", "6033", "vaccine-covid-19" ))
	if err != nil {
		log.Fatalln(err)
	}
	//defer DB.Close()
	fmt.Println("you")
	sugarLogger.Error("Test log err")
	productRepo := repository.NewProductRepository(DB)
	fmt.Println("you two")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	//products, err := productRepo.All(ctx)
	//fmt.Println(err)
	product, err := productRepo.ByID(ctx, "c370c173-19e2-11ec-a931-0242ac170002")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%+v\n", product)
	//queries := db.New(DB)
	//products, err := queries.GetProducts(ctx)
	//if err != nil {
	//	log.Println(err)
	//}
	//log.Println(products)
}
