package main

import (
	"log"
	"os"

	"github.com/ruziba3vich/armiya-gateway/api"
	authhandlers "github.com/ruziba3vich/armiya-gateway/api/handlers/auth_handlers"
	producthandlers "github.com/ruziba3vich/armiya-gateway/api/handlers/product_handlers"
	"github.com/ruziba3vich/armiya-gateway/config"
	"github.com/ruziba3vich/armiya-gateway/genprotos"
	"google.golang.org/grpc"
)

func main() {
	var cfg config.Config
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	if err := cfg.Load(); err != nil {
		logger.Fatal(err)
	}

	conn, err := grpc.Dial(cfg.AuthHost, grpc.WithInsecure())
	if err != nil {
		logger.Fatalf("Failed to connect to auth service: %v", err)
	}

	connProduct, err := grpc.Dial(cfg.ProductHost, grpc.WithInsecure())
	if err != nil {
		logger.Fatalf("Failed to connect to product service: %v", err)
	}

	authClient := genprotos.NewAuthServiceClient(conn)

	authHandlers := authhandlers.NewAuthHandlers(authClient, logger)

	productClient := genprotos.NewProductServiceClient(connProduct)

	productHandlers := producthandlers.NewProductHandlers(productClient, logger)

	api := api.New(&cfg, logger, authHandlers, productHandlers)
	logger.Fatal(api.RUN())
}
