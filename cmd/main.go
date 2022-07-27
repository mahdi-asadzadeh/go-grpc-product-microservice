package main

import (
	"fmt"
	"github.com/mahdi-asadzadeh/go-grpc-product-microservice/pkg/config"
	"github.com/mahdi-asadzadeh/go-grpc-product-microservice/pkg/db"
	"github.com/mahdi-asadzadeh/go-grpc-product-microservice/pkg/pb"
	"github.com/mahdi-asadzadeh/go-grpc-product-microservice/pkg/services"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func main() {
	fmt.Println("Starting product server ...")
	// Load config
	config.LoadSettings(true)

	// Database config
	DB_URL := os.Getenv("DB_URL")
	h := db.Init(DB_URL)

	lis, err := net.Listen("tcp", os.Getenv("SERVER_IP"))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Register product service
	grpcServer := grpc.NewServer()
	pb.RegisterProductServiceServer(grpcServer, &services.Server{H: h})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
