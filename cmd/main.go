package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/je117er/ocg-final-project/internal/product/repository"
	"log"
)

var (
	ctx context.Context
)

func main() {
	db, err := sql.Open("mysql", "admin:123456Aa@tcp(127.0.0.1:3306)/vaccine-covid-19")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
	fmt.Println("you")
	productRepo := repository.NewProductRepository(db)
	fmt.Println("you two")
	products, err := productRepo.All(ctx)
	fmt.Println("you three")
	//product, err := productRepo.ByID(ctx, "c370c173-19e2-11ec-a931-0242ac170002")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%+v\n", products)
}
