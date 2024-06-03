package main

import (
	"context"
	"cv-service/internal/config"
	"cv-service/internal/repository/mongo"
	"cv-service/internal/service"
	"cv-service/proto"
	"fmt"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	host = "0.0.0.0"
	port = "5011"
)

func main() {
	ctx := context.Background()

	err := setupViper()

	if err != nil {
		log.Fatalf("error reading yml file: %v", err)
	}

	addr := fmt.Sprintf("%s:%s", host, port)
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("error starting tcp listener: %v", err)
	}

	mongoDataBase, err := config.SetupMongoDataBase(ctx)

	if err != nil {
		log.Fatalf("error starting mongo: %v", err)
	}

	cvRepository := mongo.NewCVRepository(mongoDataBase.Collection("cv"))
	cvService := service.NewCVService(cvRepository)

	grpcServer := grpc.NewServer()

	proto.RegisterCVServiceServer(grpcServer, cvService)

	log.Printf("gRPC stated at %v\n", port)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("error statring gRPC: %v", err)
	}

}

func setupViper() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}
