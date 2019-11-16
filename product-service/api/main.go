package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/crgimenes/goconfig"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/klferreira/hash-challenge/product-service/api/handlers"
	"github.com/klferreira/hash-challenge/product-service/pkg/product/repository"
	"github.com/klferreira/hash-challenge/product-service/pkg/product/service"
	discount "github.com/klferreira/hash-challenge/product-service/proto"
	"google.golang.org/grpc"
)

type Config struct {
	DiscountServiceURI string `cfg:"DISCOUNT_SERVICE_URI" cfgDefault:"0.0.0.0:50052" cfgRequired:"true"`
	DatabaseHost       string `cfg:"DATABASE_URI" cfgDefault:"root:toor@tcp(127.0.0.1:3306)/hashchallenge" cfgRequired:"true"`
	ProductServicePort string `cfg:"PRODUCT_SERVICE_PORT" cfgDefault:":8080" cfgRequired:"true"`
}

func main() {
	r := mux.NewRouter()

	config := Config{}
	if err := goconfig.Parse(&config); err != nil {
		log.Fatal("Could not parse environment vars")
	}

	conn, err := grpc.Dial(config.DiscountServiceURI, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Could not connect to discount service")
	}

	discountClient := discount.NewDiscountServiceClient(conn)

	r.Use(handlers.AuthenticationMiddleware)
	db, err := sql.Open("mysql", config.DatabaseHost)
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewMySQLProductRepository(db)
	service := service.NewProductService(repo, discountClient)

	handlers.GetProductHandlers(r, service)

	fmt.Printf("Listening on port %s", config.ProductServicePort)
	http.ListenAndServe(config.ProductServicePort, r)
}
